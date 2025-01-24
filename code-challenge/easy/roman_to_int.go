package easy

import "fmt"

var ROMAN_TO_INT = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	res := 0

	for i, c := range s {
		if i == 0 {
			res = ROMAN_TO_INT[string(c)]
			continue
		}

		if ROMAN_TO_INT[string(c)] > ROMAN_TO_INT[string(s[i-1])] {
			res -= ROMAN_TO_INT[string(s[i-1])] * 2
		}
		res += ROMAN_TO_INT[string(c)]
	}

	return res
}

func RomanToInt(s string) {
	fmt.Println(romanToInt(s))
}
