package solution

func mergeSubResult(subResult [][]int, head int) [][]int {
	var result [][]int
	for _, nums := range subResult {
		nums = append([]int{head}, nums...)
		result = append(result, nums)
	}
	return result
}

func permute(nums []int) [][]int {
	// 最简单的三种情况
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}

	if len(nums) == 2 {
		return [][]int{
			[]int{nums[0], nums[1]},
			[]int{nums[1], nums[0]},
		}
	}

	// 长度超过2时，拆分原来的数据，迭代计算
	// 迭代计算方法：
	// 1. 将数组中的第 i+1 个数与数组中的第 1 个数互换位置
	// 2. 将数组拆分成两个部分：a：nums[0]；b：nums[1:]
	// 3. 对 nums[1:]套用此算法计算结果 subResult
	// 4. 将计算出来的 subResult 中的每一项 []int 的前面都加上原来的 nums[0]，成为此次迭代计算的最终结果
	// 5. 将此次迭代计算的最终结果合并到总的结果中
	// 6. 将数组中的第 i+1 个数与数组中的第 1 个数互换位置，还原数组

	var result [][]int
	for i := 0; i < len(nums); i++ {
		// 1
		nums[i], nums[0] = nums[0], nums[i]

		// 2
		var head = nums[0]
		var tail = nums[1:]

		// 3
		var subResult = permute(tail)

		// 4 + 5
		result = append(result, mergeSubResult(subResult, head)...)

		// 6
		nums[i], nums[0] = nums[0], nums[i]
	}

	return result
}
