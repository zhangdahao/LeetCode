func addBinary(a string, b string) string {
	if len(a) <= len(b) {
		b, a = a, b
	}
	bLen := len(b)
	diffLen := len(a) - len(b)
	var flag int
	str := ""
	for i := len(a) - 1; i >= 0; i-- {
		aNum := int(a[i] - 48)
		bNum := 0
		if i-diffLen < bLen && i >= diffLen {
			bNum = int(b[i-diffLen] - 48)
		}
		str = strconv.Itoa((aNum+bNum+flag)%2) + str
		if bNum+aNum+flag > 1 {
			flag = 1
		} else {
			flag = 0
		}
	}
	if flag == 1 {
		str = "1" + str
	}
	return str
}
