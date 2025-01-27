package attachments

import "fmt"

//Common propertity of every on-road vechicles
type Vechicle interface {
	MovesOnRoad()
}

type Bike struct{}

func (b Bike) MovesOnRoad() {
	fmt.Println("Bikes moves on the road using two wheels!")
}

type Car struct{}

func (c Car) MovesOnRoad() {
	fmt.Println("Car moves on the road using four wheels!")
}

type Auto struct{}

func (a Auto) MovesOnRoad() {
	fmt.Println("Auto moves on the road using three wheels!")
}

type Cycle struct{}

func (c Cycle) MovesOnRoad() {
	fmt.Println("Cycle moves on the road using two wheels!")
}
