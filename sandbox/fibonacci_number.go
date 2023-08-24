package main

import "fmt"

func FibonacciNumber(number int) int {
	var genereteFibonacciNumbers []int

	for i := 0; i <= number; i++ {
		genereteFibonacciNumbers = append(genereteFibonacciNumbers, i)
	}

	fmt.Printf("%v", genereteFibonacciNumbers)

	return 10
}
