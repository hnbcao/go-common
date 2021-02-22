package _03

// 效率kuai
func lengthOfLongestSubstring1(s string) int {
	var arr [128]int
	ret := 0
	lasted := 0
	for index, value := range s {
		if arr[value] > lasted {
			lasted = arr[value]
		}
		arr[value] = index + 1

		if index+1-lasted > ret {
			ret = index + 1 - lasted
		}
	}
	return ret
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
