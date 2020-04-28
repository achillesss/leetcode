package solution

import (
	"sort"
)

// 两种解法：
// 1. 合并两个有序数组，从合并后的数组中找到中位数
// 2. 通过中位数的定义，在两组数中找到合适的中位数

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

// 合并结果也为一个顺序数组
func mergeSortedArrays(nums1, nums2 []int) []int {
	var nums = append(nums1, nums2...)
	// 时间复杂度主要在合并数组之后的排序算法上: O(nlogn)
	sort.Ints(nums)
	return nums
}

func findMedianSortedArrays(nums1, nums2 []int) (float64, bool) {
	var nums = mergeSortedArrays(nums1, nums2)
	return findMedianInSortedArray(nums)
}
