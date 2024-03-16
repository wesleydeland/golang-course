package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)

	transformed := transformNumbers(&numbers, double)

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
