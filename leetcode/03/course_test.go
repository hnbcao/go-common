package _03

import (
	"fmt"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	testStr := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
		"",
	}
	for _, value := range testStr {
		fmt.Println(lengthOfLongestSubstring1(value))
		//fmt.Println(lengthOfLongestSubstring2(value))
	}
}
