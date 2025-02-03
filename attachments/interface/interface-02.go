package interfaces

import "fmt"

type Animal interface {
	MakeSound()
}

type Dog struct{}

func (d Dog) MakeSound() {
	fmt.Println("BowBow!")
}

type Cat struct{}

func (c Cat) MakeSound() {
	fmt.Println("MeowMeow!")
}

type Bird struct{}

func (b Bird) MakeSound() {
	fmt.Println("TweetTweet!")
}

func PrintInMainIf02() {
	dog := Dog{}
	cat := Cat{}
	bird := Bird{}
	dog.MakeSound()
	cat.MakeSound()
	bird.MakeSound()
}
