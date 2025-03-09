package basic

import (
	"fmt"
	"time"
)

func demoSwitch() {
	i := 2

	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		println("3")
	default:
		fmt.Println("default")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("It's the weekday")
	}

	whatIamI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Iam a bool")
		case int:
			fmt.Println("Iam an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatIamI(true)
	whatIamI(1)
	whatIamI("hey")
}
