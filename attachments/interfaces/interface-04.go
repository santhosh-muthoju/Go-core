package interfaces //polymorphism

import "fmt"

type HumanBeing interface {
	Activity()
}

type Man struct{}

type Women struct{}

type Boy struct{}

type Girl struct{}

func (m Man) Activity() {
	fmt.Println("A Man can go to gym")
}

func (w Women) Activity() {
	fmt.Println("A Women can feed her childern")
}

func (b Boy) Activity() {
	fmt.Println("A Boy goes to school")
}

func (g Girl) Activity() {
	fmt.Println("A Girl goes to school")
}

func PrintInMainIf04() {
	man := Man{}
	women := Women{}
	boy := Boy{}
	girl := Girl{}

	humans := []HumanBeing{man, women, boy, girl}

	for _, human := range humans {
		human.Activity()
	}
}
