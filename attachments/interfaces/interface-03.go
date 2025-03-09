package interfaces

import "fmt"

type Test interface {
	FinalExam()
}

type Animals interface {
	HasLegs()
}

type Dogs struct{}

func (d Dogs) HasLegs() {
	fmt.Println("A dog has four legs")
}

type Monkeys struct{}

func (m Monkeys) HasLegs() {
	fmt.Println("A monkey has two legs")
}

type Tenth struct{}

func (t Tenth) FinalExam() {
	fmt.Println("ssc")
}

type Inter struct{}

func (i Inter) FinalExam() {
	fmt.Println("12th")
}

type Graduation struct{}

func (g Graduation) FinalExam() {
	fmt.Println("Final semester Exam")
}

func PrintInMain03() {
	dog := Dogs{}
	monkey := Monkeys{}

	animals := []Animals{dog, monkey}
	for _, animal := range animals {
		animal.HasLegs()
	}

	ten := Tenth{}
	twelve := Inter{}
	grad := Graduation{}

	standard := []Test{ten, twelve, grad}
	for _, class := range standard {
		class.FinalExam()
	}
}
