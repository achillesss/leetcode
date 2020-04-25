package solution

import "fmt"

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

// 二分法
func splitArray(nums []int) ([]int, []int) {
	if len(nums) < 2 {
		return nil, nums
	}

	if len(nums)&1 == 1 {
		var index = (len(nums) - 1) / 2
		return nums[:index], nums[index:]
	}

	var index = len(nums) / 2
	return nums[:index], nums[index:]
}

// 二分法合并两个顺序数组
// 合并结果也为一个顺序数组
func mergeSortedArrays(nums1, nums2 []int) []int {
	var l1 = len(nums1)
	var l2 = len(nums2)
	if l1 == 0 {
		return nums2
	}

	if l2 == 0 {
		return nums1
	}

	var firstNums1 = nums1[0]
	var firstNums2 = nums2[0]
	var lastNums1 = nums1[l1-1]
	var lastNums2 = nums2[l2-1]

	if firstNums1 >= lastNums2 {
		return append(nums2, nums1...)
	}

	if lastNums1 <= firstNums2 {
		return append(nums1, nums2...)
	}

	if firstNums1 == firstNums2 {
		nums1 = append([]int{firstNums1}, nums1...)
		if l2 == 1 {
			nums2 = []int{}
		} else {
			nums2 = nums2[1:]
		}

		return mergeSortedArrays(nums1, nums2)
	}

	if lastNums1 == lastNums2 {
		nums1 = append(nums1, lastNums1)
		if l2 == 1 {
			nums2 = []int{}
		} else {
			nums2 = nums2[:l2-2]
		}
		return mergeSortedArrays(nums1, nums2)
	}

	if firstNums1 < firstNums2 && lastNums1 > lastNums2 {
		return mergeSortedArrays(nums2, nums1)
	}

	if firstNums1 > firstNums2 && lastNums1 < lastNums2 {
		var nums2Head, nums2Tail = splitArray(nums2)
		return mergeSortedArrays(nums2Head, mergeSortedArrays(nums1, nums2Tail))
	}

	if firstNums1 > firstNums2 && lastNums1 > lastNums2 {
		return mergeSortedArrays(nums2, nums1)
	}

	if firstNums1 < firstNums2 && lastNums1 < lastNums2 {
		var nums2Head, nums2Tail = splitArray(nums2)
		return mergeSortedArrays(nums2Head, mergeSortedArrays(nums1, nums2Tail))
	}

	panic(fmt.Sprintf("nums1: %v, nums2: %v\n", nums1, nums2))
}

func findMedianInSortedArrays(nums1, nums2 []int) (float64, bool) {
	var nums = mergeSortedArrays(nums1, nums2)
	fmt.Printf("merged: %+v\n", nums)
	return findMedianInSortedArray(nums)
}
