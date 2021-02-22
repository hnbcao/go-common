package _13

import "strings"

func romanToInt(s string) int {
	sum := 0
	romanMap := romanMap()
	s = format(s)
	for _, ch := range s {
		sum += romanMap[string(ch)]
	}
	return sum
}

func format(s string) string {
	s = strings.ReplaceAll(s, "IV", "O")
	s = strings.ReplaceAll(s, "IX", "P")
	s = strings.ReplaceAll(s, "XL", "Q")
	s = strings.ReplaceAll(s, "XC", "R")
	s = strings.ReplaceAll(s, "CD", "S")
	s = strings.ReplaceAll(s, "CM", "T")
	return s
}

func romanMap() map[string]int {
	return map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
		"O": 4,
		"P": 9,
		"Q": 40,
		"R": 90,
		"S": 400,
		"T": 900,
	}
}
