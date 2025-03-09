package functions

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func PrintClouser() {
	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
