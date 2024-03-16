package main

import "fmt"

func main() {
	age := 32 //regular variable

	agePointer := &age

	fmt.Println("Age: ", *agePointer)

	editAgeToAdultYears(agePointer)
	fmt.Println("Adult years: ", *agePointer)
}

func editAgeToAdultYears(age *int) {
	//return *age - 18
	*age = *age - 18
}
