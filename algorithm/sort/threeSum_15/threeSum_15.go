package threeSum_15

import "fmt"

func main() {

	nums := []int{8, 6, 4, 2, 3, 5, 7, 9, 1, -1, -3 - 5, -1}
	//nums := []int{1}

	fmt.Println(threeSum(nums))

}

func threeSum(nums []int) [][]int {
	n := len(nums)
	quickSort(nums, 0, n-1)
	re := make([][]int, 0)

	for one := 0; one < n; one++ {
		if nums[one] > 0 {
			break
		}
		if one > 0 && nums[one] == nums[one-1] {
			continue
		}
		three := n - 1
		for tow := one + 1; tow < n; tow++ {
			if nums[one]+nums[tow] > 0 {
				break
			}
			if tow > one+1 && nums[tow] == nums[tow-1] {
				continue
			}
			for three > tow && nums[one]+nums[tow]+nums[three] > 0 {
				three--
			}
			if tow == three {
				break
			}
			if nums[one]+nums[tow]+nums[three] == 0 {
				re = append(re, []int{nums[one], nums[tow], nums[three]})
			}
		}
	}

	return re
}

func partition(data []int, l int, h int) int {
	m := data[l]
	for l < h {
		for h > l && data[h] >= m {
			h--
		}
		if h > l {
			data[l] = data[h]
		}
		for l < h && data[l] <= m {
			l++
		}
		if l < h {
			data[h] = data[l]
		}
	}
	data[l] = m
	return l
}

//1。快排
func quickSort(data []int, l int, h int) {
	if l < h {
		mid := partition(data, l, h)
		quickSort(data, l, mid-1)
		quickSort(data, mid+1, h)
	}
}
