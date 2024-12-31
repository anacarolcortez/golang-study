package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(multiplyNumbers(&numbers, double)) //pass function as value parameter
	fmt.Println(multiplyNumbers(&numbers, triple))

	multiplierFuncion := getMultiplicationFunction(&numbers)
	multipliedNumbers := multiplyNumbers(&numbers, multiplierFuncion)
	fmt.Println(multipliedNumbers)

	//anonymous functions
	transformNumbers := multiplyNumbers(&numbers, func(number int) int {
		return number * 10
	})
	fmt.Println(transformNumbers)

	//closures
	double := createTransformer(2)
	triple := createTransformer(3)
	decimal := createTransformer(10)

	doubledNumbers := multiplyNumbers(&numbers, double)
	tripledNumbers := multiplyNumbers(&numbers, triple)
	decimalNumbers := multiplyNumbers(&numbers, decimal)

	fmt.Println(doubledNumbers)
	fmt.Println(tripledNumbers)
	fmt.Println(decimalNumbers)

	//recursions
	fmt.Println(factorial(8))

	//variadic
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sumup(10, 3, 8, 9, 2))
	fmt.Println(sum(numbers...)) //spread operator to separate items from slices

}

// functions as parameters
func multiplyNumbers(numbers *[]int, multiply func(int) int) []int {
	var multipliedNumbers []int
	for _, number := range *numbers {
		multipliedNumbers = append(multipliedNumbers, multiply(number))
	}
	return multipliedNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

// functions as return values
func getMultiplicationFunction(numbers *[]int) func(int) int {
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}

// closures: capture and retain variables from the scope where they were defined.
// They allow a function to access and manipulate variables outside its own scope.
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

// recursions
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// variadic functions
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func sumup(initialValue int, numbers ...int) int {
	total := initialValue
	for _, number := range numbers {
		total += number
	}
	return total
}
