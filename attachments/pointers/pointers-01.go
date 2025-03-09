package pointers

import "fmt"

type User struct {
	email    string
	userName string
	age      int
}

func (u User) Email() string {
	return u.email
}

// type User struct {
// 	email    string
// 	userName string
// 	age      int
// }

// func (u User) Email() string {
// 	return u.email
// }

func (u *User) UpdateEmail(email string) {
	u.email = email
}

func printinmain() {
	user := User{
		email: "santhoshmuthoju@gamil.com",
	}
	user.UpdateEmail("sarayu@gmail.com")
	fmt.Println(user.Email())
}
