// 声明包文件
package main

import "fmt"

func maxSubArray(num []int) int {
	var max = num[0]
	var flag = max
	if max < 0 {
		flag = 0
	}
	for i := 1; i < len(num); i++ {
		flag = flag + num[i]
		if flag < 0 {
			if max < num[i] {
				max = num[i]
			}
			flag = 0
		} else {
			if flag > max {
				max = flag
			}
		}
	}
	return max
}
func main() {
	var nums = []int{-1, 1, 2, 1}
	fmt.Println(maxSubArray(nums))
}
