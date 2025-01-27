package attachments

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
