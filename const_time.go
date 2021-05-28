package main

import (
	"fmt"
	"unicode/utf8"
)

func CompareStrings(str1, str2 string) bool {
	var (
		res     int
		HelpIdx int
	)
	indexL1 := utf8.RuneCountInString(str1) - 1 // считаем количество символов в рунах
	indexL2 := utf8.RuneCountInString(str2) - 1

	if str1 == "" || str2 == "" {
		return false
	}

	for {
		res |= int(str1[indexL1] ^ str2[indexL2])

		if indexL1 == 0 {
			break
		}

		indexL1--

		if indexL2 != 0 {
			indexL2--
		}
		if indexL2 == 0 {
			HelpIdx++
		}
	}

	return (res + HelpIdx) == 0 //находим сколько 0 в результате
}

func main() {
	str1 := "travac" //строка злоумышленника
	str2 := "travac" // строка проверяемая на соответствие
	fmt.Println(CompareStrings(str1, str2))
}
