package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(sumUp(numbers...))
}

func sumUp(numbers ...int) int {
	sum := 0

	for _, v := range numbers {
		sum += v
	}

	return sum
}
