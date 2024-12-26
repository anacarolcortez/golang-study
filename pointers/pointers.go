package main

import "fmt"

func main() {
	age := 39
	agePointer := &age //& address -> type *int

	fmt.Println("You've been an adult for", getAdultYears(agePointer), "years")

	editAgeToAdultYears(agePointer)
	fmt.Println(age)
}

func getAdultYears(age *int) int {
	return *age - 18 //value in the address
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18 //changes de value in the address
}
