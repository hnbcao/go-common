package _07

import "math"

func reverse(x int) int {
	sum := 0
	for x != 0 {
		sum = sum*10 + (x % 10)
		x = x / 10
	}
	if math.Abs(float64(sum)) > math.Pow(2, 31) {
		return 0
	}
	return sum
}
