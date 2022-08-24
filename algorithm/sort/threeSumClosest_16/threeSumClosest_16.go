package threeSumClosest_16

import (
	"fmt"
	"math"
)

func main() {

	nums := []int{-1, 2, 1, -4}
	//nums := []int{8, 6, 4, 2, 3, 5, 7, 9, 1, -1, -3 - 5, -1}
	//nums := []int{1}

	fmt.Println(threeSumClosest(nums, 1))

}

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	quickSort(nums, 0, n-1)
	re := math.MaxInt32

	for one := 0; one < n; one++ {
		if one > 0 && nums[one] == nums[one-1] {
			continue
		}

		tow := one + 1
		three := n - 1

		for tow < three {

			sum := nums[one] + nums[tow] + nums[three]
			if sum == target {
				return target
			}

			re = updateSum(sum, re, target)

			if sum > target {
				for tow < three && nums[three] == nums[three-1] {
					three--
				}
				three--
			} else {
				for tow < three && nums[tow] == nums[tow+1] {
					tow++
				}
				tow++
			}
		}
	}

	return re
}

func updateSum(sum, re, target int) int {
	if abs(sum-target) < abs(re-target) {
		re = sum
	}
	return re
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
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
