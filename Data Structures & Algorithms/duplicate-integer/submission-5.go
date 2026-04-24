func hasDuplicate(nums []int) bool {
	duplicated := make(map[int]bool)
	for _, num := range nums {
		if _, dp := duplicated[num]; dp {
			return true
		}

		duplicated[num] = true
	}

	return false
}
