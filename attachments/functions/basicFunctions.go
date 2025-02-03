package functions

import "fmt"

func ShareCandies(friends ...string) {
	fmt.Println("Sharing candies with:")
	for _, friend := range friends {
		fmt.Println(friend)
	}
}
