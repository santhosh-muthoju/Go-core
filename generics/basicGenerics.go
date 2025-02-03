package generics

import "fmt"

// Generics is used to eliminate the necessacity of writing the repeting or duplicate code or functions

type Number interface {
	int | int32 | int64 | float32 | float64
}

//The Number interface is implemented in square [] with a variable T for generics implementation
func SumNumbers[T Number](numbers []T) T {
	var result T
	for i := range numbers {
		result += numbers[i]
	}
	return result
}

func PrintInMainGs01() {
	numbers1 := []int{1, 2, 3, 4, 5}
	numbers2 := []int32{1, 2, 3, 4, 5}
	numbers3 := []int64{1, 2, 3, 4, 5}
	numbers4 := []float32{1.1, 2.1, 3.1, 4.1, 5.1}
	numbers5 := []float64{1.1, 2.1, 3.1, 4.1, 5.1}

	fmt.Println("sum of numbers1:", SumNumbers(numbers1))
	fmt.Println("sum of numbers1:", SumNumbers(numbers2))
	fmt.Println("sum of numbers1:", SumNumbers(numbers3))
	fmt.Println("sum of numbers1:", SumNumbers(numbers4))
	fmt.Println("sum of numbers1:", SumNumbers(numbers5))
}
