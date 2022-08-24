package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 4, 6, 8, 9}
	//nums := []int{0,0,1,1,1,2,2,2,3,4,6}
	//nums := []int{1,1,2}
	target := searchTarget(nums, 8)
	fmt.Println(target)

}

//二分查找704
func searchTarget(nums []int, target int) int {
	//sort.SearchInts(nums,target)
	n := len(nums)
	left, right := 0, n-1
	had := -1
	for left <= right {
		mid := (right-left)/2 + left
		if target == nums[mid] {
			had = mid
			break
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return had
}

//二分查找34
func searchRange(nums []int, target int) []int {
	//sort.SearchInts(nums,target)
	n := len(nums)
	left, right := 0, n-1
	had := -1
	for left <= right {
		mid := (right-left)/2 + left
		if target == nums[mid] {
			had = mid
			break
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if had == -1 {
		return []int{-1, -1}
	}
	result := []int{had, had}

	lhad := had - 1
	rhad := had + 1
	for lhad >= 0 {
		if nums[lhad] == target {
			result[0] = lhad
			lhad--
		} else {
			break
		}
	}
	for rhad < n {
		if nums[rhad] == target {
			result[1] = rhad
			rhad++
		} else {
			break
		}
	}
	return result
}

//二分查找35
func searchInsert2(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := (right-left)/2 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

//二分查找
func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	return search(nums, l, r, target)
}

//二分查找
func search(nums []int, l, r, target int) int {
	index := 0
	m := (l + r) / 2
	if nums[m] == target {
		index = m
	}
	if nums[m] < target {
		if m >= r {
			return r + 1
		}
		return search(nums, m+1, r, target)
	}
	if nums[m] > target {
		if m == 0 {
			return m
		}
		if m <= l {
			return l
		}
		return search(nums, l, m-1, target)
	}
	return index
}


