package main

import "fmt"

// https://leetcode.com/problems/roman-to-integer/description/
func romanToInt(s string) int {
	m := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	r := []rune(s)

	t := 0
	p := 0

	for _, k := range r {
		fmt.Println("시작:", string(k), m[k])
		if p == 0 {
			p = m[k]
			fmt.Println("p가 0일 때:", p, m[k])
			continue
		}

		if p >= m[k] {
			t += p
			p = m[k]
			fmt.Println("p가 m[k]보다 같거나 클 때:", p, m[k])
			continue
		} else {
			p = m[k] - p
			fmt.Println("p가 m[k]보다 적을 때:", p, m[k])
		}

	}
	t += p
	return t
}

func main() {
	fmt.Printf("result: %d\n", romanToInt("MCMXCIV"))
}
