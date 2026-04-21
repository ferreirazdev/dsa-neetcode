func twoSum(nums []int, target int) []int {
    compare_list := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if j, found := compare_list[diff]; found {
			return []int{j, i}
		}

		compare_list[num] = i
	}

	return []int{}
}
