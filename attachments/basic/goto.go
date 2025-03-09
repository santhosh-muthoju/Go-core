package basic

import "fmt"

func Printgoto() {
	count := 1

start:
	fmt.Println("count:", count)
	count++

	if count <= 5 {
		goto start
	}

	fmt.Println("Done!")
}
