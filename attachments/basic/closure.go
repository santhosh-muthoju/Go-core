package basic

import "fmt"

var i int = 0

// func keyword will be 2 times in a closure as it returns a function
func addNum() func() int {
	//var i int
	return func() int {
		i++
		return i
	}
}

func PrintInMainClosure() {
	newFunc := addNum()
	fmt.Println("value 1 :", newFunc())
	fmt.Println("value 2 :", newFunc())
	fmt.Println("value 3 :", newFunc())

	newFunc2 := addNum()

	fmt.Println("new value :", newFunc2())
}
