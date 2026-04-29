func isValid(s string) bool {
    stack := []rune{}
	closeToOpen := map[rune]rune{
		')':'(', 
		']':'[', 
		'}':'{',
	}

	for _, c := range s {
		if open, valid := closeToOpen[c]; valid {
			if len(stack) > 0 && stack[len(stack)-1] == open {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, c)
		}
	}

	return len(stack) == 0
}