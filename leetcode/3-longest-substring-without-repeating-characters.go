package main

func lengthOfLongestSubstring(s string) int {

	runes := []rune(s)
	list := make(map[rune]int)
	left := 0
	currentLength := 0
	maxLength := 0
	for right := range runes {

		if index, exists := list[runes[right]]; exists && index >= left {
			left = index + 1
			list[runes[right]] = left
		}
		list[runes[right]] = right
		currentLength = right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}
