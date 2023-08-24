package main

import "fmt"

func PrimeNumber(number int) int {
	var generetePrimeNumber []int

	for i := 2; i <= number; i++ {
		for j := 2; j <= i; j++ {
			if i == j { // для 2
				generetePrimeNumber = append(generetePrimeNumber, i)
			}

			if i%j == 0 {
				break
			}

			generetePrimeNumber = append(generetePrimeNumber, i)
			break
		}
	}

	fmt.Printf("%v", generetePrimeNumber)

	return 10
}
