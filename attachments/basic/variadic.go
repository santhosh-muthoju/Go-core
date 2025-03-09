package basic

import "fmt"

func PrintInMainVariadic() {
	n1 := []int{1, 2, 3, 4, 5}
	demoVariadic(n1...)
}

func demoVariadic(nums ...int) {
	nums = append(nums, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	var sum int
	for _, value := range nums {
		sum += value
	}
	fmt.Print(sum)
}
