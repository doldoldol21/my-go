package main

import "fmt"

import "strconv"

// 11ms
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	reverse := string(runes)
	return s == reverse
}

func main() {
	fmt.Println(isPalindrome(10))   // false
	fmt.Println(isPalindrome(-121)) // false
	fmt.Println(isPalindrome(121))  // true
}
