package main

func HappyTicket(number int) bool {
	var numberSlice []int
	for number > 0 {
		numberSlice = append(numberSlice, number%10)
		number /= 10
	}

	var sumLeftPart = 0
	for i := 0; i < len(numberSlice)/2; i++ {
		sumLeftPart += numberSlice[i]
	}

	var sumRightPart = 0
	if len(numberSlice)%2 != 0 { // нечетное
		for i := len(numberSlice)/2 + 1; i < len(numberSlice); i++ {
			sumRightPart += numberSlice[i]
		}
	} else { // четное
		for i := len(numberSlice) / 2; i < len(numberSlice); i++ {
			sumRightPart += numberSlice[i]
		}
	}

	if sumLeftPart != sumRightPart {
		return false
	}

	return true
}
