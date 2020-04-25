package solution

import (
	"fmt"
	"sort"
)

// 从一个有序数组中找到中位数
func findMedianInSortedArray(nums []int) (float64, bool) {
	if nums == nil {
		return 0, false
	}

	if len(nums)&1 == 1 {
		return float64(nums[((len(nums) - 1) / 2)]), true
	}

	return (float64(nums[len(nums)/2]) + float64(nums[len(nums)/2-1])) / 2, true
}

// 二分法合并两个顺序数组
// 合并结果也为一个顺序数组
func mergeSortedArrays(nums1, nums2 []int) []int {
	var nums = append(nums1, nums2...)
	sort.Ints(nums)
	return nums
}

func findMedianSortedArrays(nums1, nums2 []int) (float64, bool) {
	var nums = mergeSortedArrays(nums1, nums2)
	fmt.Printf("merged: %+v\n", nums)
	return findMedianInSortedArray(nums)
}
