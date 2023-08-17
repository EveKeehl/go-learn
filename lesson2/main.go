package main

import (
	"fmt"
)

//func main() {
//	result := itoa(12_345)
//	fmt.Print(result)
//}

func main() {
	type pair struct {
		i int
		s string
	}

	test := []pair{
		{0, "0"},
		{22, "22"},
		{32432523, "32432523"},
		{-3, "-3"},
	}
	for _, t := range test {
		if t.s == itoa(t.i) {
			fmt.Printf("%d - %s\n", t.i, "OK")
		} else {
			fmt.Printf("%d - %s\n", t.i, "FAIL")
		}
	}
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

		numberString = string('0'+digit) + numberString

		number /= 10
	}

	if isNegative {
		numberString = "-" + numberString
	}

	return numberString
}
