package _1

func twoSum(nums []int, target int) []int {
	for i, num1 := range nums {
		tmp := nums[i+1 : len(nums)]
		for j, num2 := range tmp {
			if target == num1+num2 {
				return []int{i, j + i + 1}
			}
		}
	}
	return nil
}
