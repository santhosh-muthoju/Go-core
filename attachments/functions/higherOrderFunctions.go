package functions

import "fmt"

//passing function as a parameter

func applyOperations(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func PrintHOF() {
	add := func(x, y int) int {
		return x + y
	}
	result := applyOperations(10, 20, add)
	fmt.Println(result)
}

func multiply(a, b int, mul func(int, int) int) int {
	return mul(a, b)
}

func PrintHOF2() {
	mul := func(x, y int) int {
		return x * y
	}
	res := multiply(10, 10, mul)
	fmt.Println(res)
}

//returning a function a return type
