package _29

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	matrix := [][]int{
		{2, 3}, //1,0 = 8 1,1 = 9 1,2 = 4

	}
	res := spiralOrder(matrix)
	fmt.Println(res)
}
