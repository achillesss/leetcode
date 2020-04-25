package solution

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	var nums = []int{1}
	var s = permute(nums)
	fmt.Println(nums, len(s), " results: ", s)
	nums = []int{1, 2}
	s = permute(nums)
	fmt.Println(nums, len(s), " results: ", s)
	nums = []int{1, 2, 3}
	s = permute(nums)
	fmt.Println(nums, len(s), " results: ", s)
	nums = []int{1, 2, 3, 4}
	s = permute(nums)
	fmt.Println(nums, len(s), " results: ", s)
	nums = []int{1, 2, 3, 4, 5}
	s = permute(nums)
	fmt.Println(nums, len(s), " results: ", s)
}
