package main

import "errors"

/*
Распаковка строки
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:

* "a4bc2d5e" => "aaaabccddddde"
* "abcd" => "abcd"
* "45" => "" (некорректная строка)

Дополнительное задание: поддержка escape - последовательности
* `qwe\4\5` => `qwe45` (*)
* `qwe\45` => `qwe44444` (*)
* `qwe\\5` => `qwe\\\\\` (*)
*/

var ErrIncorrectString = errors.New("incorrect string")

func UnpackingString(string string) (string, error) {
	return string, ErrIncorrectString
}
