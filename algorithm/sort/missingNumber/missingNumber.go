package main

import "fmt"

func main() {

	nums := []int{8, 6, 4, 2, 3, 5, 7, 9, 1}
	//nums := []int{1}

	fmt.Println(missingNumber(nums))
}

func missingNumber(nums []int) int {
	n := len(nums)
	quickSort(nums, 0, n-1)
	i := n
	for ; i > 0; i-- {
		if nums[i-1] != i {
			return i
		}
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
