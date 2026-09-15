package main

import "fmt"
import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := append(nums1, nums2...)
	sort.Ints(merged)
	fmt.Println(merged)
	isEven := len(merged)%2 == 0
	if !isEven {
		return float64(merged[len(merged)/2])
	}
	result := (float64(merged[len(merged)/2-1]) + float64(merged[len(merged)/2])) / 2
	return result
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))    // 2
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4})) // 2.5
}
