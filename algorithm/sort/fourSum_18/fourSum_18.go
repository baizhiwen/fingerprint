package main

import "fmt"

func main() {

	nums := []int{2,2,2,2,2}
	target := 8

	fmt.Println(fourSum(nums,target))

}

func fourSum(nums []int,target int) [][]int {
	n := len(nums)
	quickSort(nums, 0, n-1)
	re := make([][]int, 0)

	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3]<=target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-1]+nums[n-2]+nums[n-3]<target{
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-1]+nums[n-2] <target {
				continue
			}
			for l,r := j+1,n-1;l<r; {
				if sum :=nums[i]+nums[j]+nums[l]+nums[r];sum == target {
					re = append(re, []int{nums[i], nums[j], nums[l], nums[r]})
					for l++; l < r && nums[l] == nums[l-1]; l++ {}
					for r--; l < r && nums[r] == nums[r+1]; r-- {}
				}else if sum < target {
					l++
				}else {
					r--
				}
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
