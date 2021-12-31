package main

import "fmt"

func main() {

	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	fmt.Println(merge(nums1, 3, nums2, 3))
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {

	//if m == 0 {
	//	nums1 = nums2
	//}
	//if n == 0 {
	//	nums1 = nums1[:m]
	//}
	sum := make([]int, 0, m+n)
	var i, j int
	for {
		if i == m {
			sum = append(sum, nums2[j:]...)
			break
		}
		if j == n {
			sum = append(sum, nums1[i:]...)
			break
		}

		if nums1[i] < nums2[j] {
			sum = append(sum, nums1[i])
			i++
		} else if nums1[i] > nums2[j] {
			sum = append(sum, nums2[j])
			j++
		} else {
			sum = append(sum, nums1[i])
			i++
			sum = append(sum, nums2[j])
			j++
		}

	}

	copy(nums1, sum)

	return nums1
}
