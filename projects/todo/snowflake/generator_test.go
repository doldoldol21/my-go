package snowflake

import (
	"os"
	"testing"
)

// ID 중복 및 유효성 검증
func TestSnowflakeGenerator(t *testing.T) {
	ids := make(map[int64]bool)

	for range 100 {
		id, err := NextId()
		if err != nil {
			t.Fatalf("ID 생성 실패: %v", err)
		}
		if ids[id] {
			t.Errorf("중복 ID 발견: %d", id)
		}
		ids[id] = true

		if id < 0 {
			t.Errorf("잘못된 ID: %d (양수여야 함)", id)
		}
	}
}

// MAC 주소 기반 머신 ID 생성 테스트
func TestMachineID(t *testing.T) {
	if _, err := getMachineID(); err != nil {
		t.Errorf("장치 ID 생성 실패: %v", err)
	}
}

// 환경변수 MACHINE_ID 우선 사용 테스트
func TestMachineIDFromEnv(t *testing.T) {
	os.Setenv("MACHINE_ID", "123")
	defer os.Unsetenv("MACHINE_ID")
	if _, err := getMachineID(); err != nil {
		t.Errorf("환경변수 설정 후 장치 ID 생성 실패: %v", err)
	}
}
