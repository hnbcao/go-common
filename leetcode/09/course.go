package _09

import "math"

func isPalindrome(x int) bool {
	if float64(x) < 0 || float64(x) > math.Pow(2, 31)-1 {
		return false
	}
	data := x
	res := 0
	for x != 0 {
		res = res*10 + (x % 10)
		x = x / 10
	}
	return res == data
}
