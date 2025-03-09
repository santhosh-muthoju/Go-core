package functions

import "fmt"

func PrintAnynomousFunc() {
	//Anonymous (Lambda) Functions:
	add := func(a, b int) int {
		return a + b
	}

	result := add(3, 4)
	fmt.Println(result)

	//// Immediately invoked functions:
	//ex 2:
	func(c, d int) {
		sum := c + d
		fmt.Println(sum)
	}(1, 2)

	//ex 3:
	func(name string) {
		fmt.Println("Hello,", name)
	}("Santhosh!")
}
