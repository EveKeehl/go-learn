package main

import "fmt"

func main() {
	fmt.Printf("%s", itoa(10))
}

func itoa(number int) string {
	if number == 0 {
		return "0"
	}

	var isNegative bool
	if number < 0 {
		isNegative = true
		number *= -1
	}

	var numberString string

	for number > 0 {
		digit := number % 10

		numberString = string(rune('0'+digit)) + numberString

		number /= 10
	}

	if isNegative {
		numberString = "-" + numberString
	}

	return numberString
}
