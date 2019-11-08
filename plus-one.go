//加一
func plusOne(digits []int) []int {
	flag := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if flag != 0 {
			if digits[i] == 9 {
				flag = 1
				digits[i] = 0
			} else {
				flag = 0
				digits[i]++
			}
			if flag != 0 && i == 0 {
				digits = append([]int{1}, digits...)
			}
		}
	}
	return digits
}
