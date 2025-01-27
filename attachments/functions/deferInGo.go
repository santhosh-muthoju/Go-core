package attachments

import "fmt"

func PrintForSampleDefer() {
	fmt.Println("Start Playing...")
	defer fmt.Println("Keep all the toys in the box!")
	fmt.Println("continuing playing...")
	fmt.Println("finished playing...!")
}
