package pointers

import "fmt"

func DemoPointer() {
	var num1 int = 100
	var num2 *int = &num1 // As num2 is a pointer to num1 this should store it's value as address of num1

	//now num2 is a pointer to num1

	// lets print the address of num1
	fmt.Println("The address of num1:", &num1)

	// lets print value assigned to num2
	fmt.Println("The value assigned to num2:", num2)

	// You can see in the O/P that both the above values are same

	//Now '*' can be used to access the value of num1 by placing it infront of its pointer which is num2
	//This process of acessing the value of original var through its pointer is called "dereferencing".
	fmt.Println(*num2)

	// The above O/P is equal to the value of num which is 100 :)

	// Now lets try changing the value of num1 using its pointer num2 and print num1 value
	*num2 = 200
	fmt.Println("The value of num1:", num1)

	//For the above O/P you see num1 = 200 means a pointer can be used to
	// change the value of original variable where its pointing to using '*' :)!

}
