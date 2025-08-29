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