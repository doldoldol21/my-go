package snowflake

import (
	"fmt"
	"hash/fnv"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	instance *SnowflakeGenerator
	once     sync.Once
)

// Snowflake ID 생성기 (분산 환경에서 고유 ID 생성)
type SnowflakeGenerator struct {
	machineID int64 // 머신 식별자 (0-1023)
	sequence  int64 // 시퀀스 번호 (같은 밀리초 내 중복 방지)
	lastTime  int64 // 마지막 생성 시간
	mutex     sync.Mutex
}

func NextId() (int64, error) {
	generator := getSnowflakeGenerator()
	generator.mutex.Lock()
	defer generator.mutex.Unlock()
	now := time.Now().UnixMilli()
	// 시계 역행 감지
	if now < generator.lastTime {
		return 0, fmt.Errorf("clock moved backwards by %d ms", generator.lastTime-now)
	} else if now == generator.lastTime {
		// 같은 밀리초 내 시퀀스 증가
		generator.sequence = (generator.sequence + 1) & 0xFFF
		if generator.sequence == 0 {
			// 같은 밀리초에 4096개 초과 시 대기
			for now <= generator.lastTime {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		generator.sequence = 0
	}
	generator.lastTime = now
	// Snowflake ID 생성: 타임스탬프(41bit) + 머신ID(10bit) + 시퀀스(12bit)
	return ((now - 1756252800000) << 22) | (generator.machineID << 12) | generator.sequence, nil
}

func getSnowflakeGenerator() *SnowflakeGenerator {
	once.Do(func() {
		machineID, err := getMachineID()
		if err != nil {
			panic("can't found machineID")
		}
		if machineID < 0 || machineID > 1023 {
			panic("machineID must be between 0 and 1023")
		}
		instance = &SnowflakeGenerator{
			machineID: machineID,
			sequence:  0,
			lastTime:  0,
		}
	})
	return instance
}

// 머신 ID 생성: 환경변수 우선, 없으면 MAC 주소 해시 사용
func getMachineID() (int64, error) {
	if id := os.Getenv("MACHINE_ID"); id != "" {
		machineID, _ := strconv.ParseInt(id, 10, 64)
		return machineID % 1024, nil
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		return -1, fmt.Errorf("error: %s", err)
	}
	for _, inter := range interfaces {
		if inter.HardwareAddr != nil {
			macAddr := inter.HardwareAddr.String()
			h := fnv.New32a()
			h.Write([]byte(macAddr))
			hash := h.Sum32()
			return int64(hash & 1023), nil
		}
	}
	return -1, fmt.Errorf("error: not found hardware address. %v", err)
}
