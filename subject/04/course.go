package _04

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1)+len(nums2) == 0 {
		return 0
	}
	sum := 0.0
	nums1Index, nums2Index := 0, 0
	mid := (len(nums1) + len(nums2)) / 2
	isTwo := (len(nums1)+len(nums2))%2 == 0
	for i := 0; i <= mid; i++ {
		current := 0
		if nums1Index >= len(nums1) {
			current = nums2[nums2Index]
			nums2Index++
		} else if nums2Index >= len(nums2) {
			current = nums1[nums1Index]
			nums1Index++
		} else {
			if nums1[nums1Index] > nums2[nums2Index] {
				current = nums2[nums2Index]
				nums2Index++
			} else {
				current = nums1[nums1Index]
				nums1Index++
			}
		}
		if i == mid-1 && isTwo {
			sum += float64(current)
		}
		if i == mid {
			sum += float64(current)
		}
	}
	if isTwo {
		sum = sum / 2
	}
	return sum
}
