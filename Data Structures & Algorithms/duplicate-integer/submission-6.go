func hasDuplicate(nums []int) bool {
	prevMap := make(map[int]bool)
	for _, num := range nums {
		if _, duplicated := prevMap[num]; duplicated {
			return true
		}

		prevMap[num] = false
	}

	return false
}
