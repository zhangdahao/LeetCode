// 最后一个单词长度
func lengthOfLastWord(s string) int {
	var num, flag = 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			if flag == 1 {
				num = 1
			} else {
				num++
			}
			flag = 0
		} else {
			flag = 1
		}
	}
	return num
}
