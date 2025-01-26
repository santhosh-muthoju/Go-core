package main

import (
	"fmt"
	"interviewPrep/internal"
)

func main() {
	user := internal.CreateUser()
	fmt.Println("User created:")
	fmt.Printf("ID: %v\n", user.ID)
	fmt.Printf("Name: %v\n", user.UserName)
	fmt.Printf("Age: %v\n", user.Age)
	fmt.Printf("Gender: %s\n", user.Gender)
	fmt.Printf("Profession: %s\n", user.Profession)
}
