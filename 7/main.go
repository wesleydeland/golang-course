package main

import "fmt"

type floatMap map[string]float64

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Max"
	userNames[1] = "Wesley"

	fmt.Println(userNames)

	userNames = append(userNames, "Levi")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.3

	fmt.Println(courseRatings)

	for index, value := range userNames {
		fmt.Println(index, value)
	}

	for k, v := range courseRatings {
		fmt.Println(k, v)
	}
}
