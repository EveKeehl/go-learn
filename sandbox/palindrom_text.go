package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func PalindromText(text string) bool {
	textNoSpaces := strings.ReplaceAll(text, " ", "")

	textRune := []rune(textNoSpaces)
	length := len(textRune)

	for i := 0; i < length/2; i++ {
		if textRune[i] != textRune[length-1-i] {
			return false
		}
	}

	return true
}

func PalindromTextMY(text string) string {
	textNoSpaces := strings.ReplaceAll(text, " ", "")
	lengthString := utf8.RuneCountInString(textNoSpaces) // получить длинну строки в рунах = символах

	var part1 []rune
	var part2 []rune

	textRune := []rune(textNoSpaces)

	part1 = textRune[:lengthString/2] // слайс рун с 0 символа до середины

	if lengthString%2 == 0 {
		part2 = textRune[lengthString/2:] // слайс рун с середины до конца
	} else {
		part2 = textRune[lengthString/2+1:] // слайс рун с середины + 1 до конца
	}

	part2Reversed := []rune{} // развернули part2 для удобства сравнения

	for i := len(part2) - 1; i >= 0; i-- {
		part2Reversed = append(part2Reversed, part2[i])
	}

	result := true // флаг для сравнения слайсов

	for i, v := range part1 { // сравнение слайсов
		if v != part2Reversed[i] {
			result = false
		}
	}

	if result {
		return fmt.Sprintf("\"%s\" палиндром", text)
	}

	return " \"" + text + "\" " + "не палиндром \n"
}
