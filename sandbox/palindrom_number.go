package main

func PalindromNumber(number int) bool {
	var numberSlice []int
	for number > 0 {
		numberSlice = append(numberSlice, number%10)
		number /= 10
	}

	length := len(numberSlice)
	for i := 0; i < length; i++ {
		if numberSlice[i] != numberSlice[length-1-i] {
			return false
		}
	}

	return true
}
