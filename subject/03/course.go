package _03

import (
	"math"
	"strings"
)

// 效率慢
func lengthOfLongestSubstring1(s string) int {
	max := 0
	length := len(s)
	for prev := 0; prev < length; prev++ {
		if length-prev < max {
			return max
		}
		str := ""
		for next := prev; next < length; next++ {
			if strings.Contains(str, string(s[next])) {
				break
			} else {
				str = str + string(s[next])
				max = int(math.Max(float64(max), float64(next-prev+1)))
				//fmt.Println(str)
			}

		}
	}
	return max
}

// 效率快
func lengthOfLongestSubstring2(s string) int {
	fast := 0
	slow := 0
	max := 0
	bitmap := [16]byte{}
	for ; fast != len(s); fast++ {
		if bitmap[s[fast]/8]&(1<<(s[fast]%8)) != 0 {
			if fast-slow > max {
				max = fast - slow
			}
			for ; s[slow] != s[fast]; slow++ {
				bitmap[s[slow]/8] ^= 1 << (s[slow] % 8)
			}
			slow++
		}
		bitmap[s[fast]/8] |= 1 << (s[fast] % 8)
	}
	if fast-slow > max {
		max = fast - slow
	}
	return max
}
