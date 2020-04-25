package solution

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	var nums1 = []int{1, 2, 100}
	var nums2 = []int{3, 4, 56, 101}
	var midian, ok = findMedianInSortedArrays(nums1, nums2)
	fmt.Printf("Midian: %v, ok: %t\n", midian, ok)
}
