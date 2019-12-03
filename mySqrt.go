func mySqrt(x int) int {

	if x < 0 {
		panic("x不能小于0")
	}

	min := 0
	max := x
	num := 0
	for true {
		if max%2 == 1 {
			max++
		}

		if min%2 == 1 {
			min--
		}

		num = (min + max) >> 1
		if num*num == x {
			break
		} else if num*num > x {
			max = num
		} else {
			min = num
		}
		if min+1 == max {
			num = min
			break
		}

	}
	return num
}
