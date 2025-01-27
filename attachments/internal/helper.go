package attachments

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}

func CreateUser() User {

	var testUser User
	testUser.ID = GenerateUUID()

	reader := bufio.NewReader(os.Stdin)

	//Get user name
	fmt.Print("Enter the user name : ")
	name, _ := reader.ReadString('\n')
	testUser.UserName = strings.TrimSpace(name)

	//Get user age
	fmt.Print("Enter the user age : ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Unable to read the input age", err)
		os.Exit(1)
	}
	input = strings.TrimSpace(input)
	age, err := strconv.Atoi(input)
	if err != nil || age < 0 {
		fmt.Println("Invalid input, please enter a valid age:", err)
	}
	testUser.Age = uint(age)

	//Get user gender
	fmt.Print("Enter the user gender : ")
	gender, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input, please enter a valid gender:", err)
		os.Exit(1)
	}
	testUser.Gender = strings.TrimSpace(gender)

	//Get user profession
	fmt.Print("Enter the user profession : ")
	prof, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input, please enter a valid profession:", err)
		os.Exit(1)
	}
	testUser.Profession = strings.TrimSpace(prof)

	return testUser
}

func PrintingData(user User) {
	fmt.Println("User created:")
	fmt.Printf("ID: %v\n", user.ID)
	fmt.Printf("Name: %v\n", user.UserName)
	fmt.Printf("Age: %v\n", user.Age)
	fmt.Printf("Gender: %s\n", user.Gender)
	fmt.Printf("Profession: %s\n", user.Profession)
}
