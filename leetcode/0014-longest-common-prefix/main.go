package main

import "fmt"

// https://leetcode.com/problems/longest-common-prefix/description/
func longestCommonPrefix(strs []string) string {
	lcp := strs[0]

	for i := 1; i < len(strs); i++ {
		j := min(len(lcp), len(strs[i]))
		for ; j > 0; j-- {
			if lcp[:j] == strs[i][:j] {
				lcp = lcp[:j]
				break
			}
		}
		if lcp[:j] == "" {
			return ""
		}
	}

	return lcp
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flight", "fliy", "fligon"})) // fli
}
