package _02

import "math/big"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	num1 := big.NewInt(0)
	num2 := big.NewInt(0)

	for i := 0; l1 != nil; i++ {
		tmp := big.NewInt(1)
		if i > 0 {
			for j := 1; j <= i; j++ {
				tmp = tmp.Mul(tmp, big.NewInt(10))
			}
		}
		num1 = num1.Add(num1, tmp.Mul(tmp, big.NewInt(int64(l1.Val))))
		l1 = l1.Next
	}
	for i := 0; l2 != nil; i++ {
		tmp := big.NewInt(1)
		if i > 0 {
			for j := 1; j <= i; j++ {
				tmp = tmp.Mul(tmp, big.NewInt(10))
			}
		}
		num2 = num2.Add(num2, tmp.Mul(tmp, big.NewInt(int64(l2.Val))))
		l2 = l2.Next
	}
	sumMod := big.NewInt(0)
	sumMod = sumMod.Add(num1, num2)
	sumDiv := big.NewInt(0)
	sumDiv = sumDiv.Add(num1, num2)

	var res *ListNode
	var last *ListNode
	for sumDiv.Cmp(big.NewInt(0)) == 1 {
		tmp := sumMod.Mod(sumDiv, big.NewInt(10))
		sumDiv = sumDiv.Div(sumDiv, big.NewInt(10))
		node := &ListNode{
			Val:  int(tmp.Int64()),
			Next: nil,
		}
		if res == nil {
			res = node
		} else {
			last.Next = node
		}
		last = node
	}
	if res == nil {
		res = &ListNode{
			Val:  0,
			Next: nil,
		}
	}
	return res
}
