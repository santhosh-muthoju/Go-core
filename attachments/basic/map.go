package basic

import (
	"fmt"
	"maps"
)

func demoMap() {
	var m = make(map[string]int)
	m["usa"] = 1
	m["china"] = 2

	fmt.Println("map:", m)

	delete(m, "china")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	newMap1 := map[string]int{"santhosh": 9154, "sarayu": 7095}
	for key, value := range newMap1 {
		fmt.Println(key, ":", value)
	}

	newMap2 := map[string]int{"santhosh": 9154, "sarayu": 7095}
	if maps.Equal(newMap1, newMap2) {
		fmt.Println("newMap1 == newMap2")
	} else {
		fmt.Println("currently not equal")
	}
}
