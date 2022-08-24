package main

import "fmt"

func main() {
	k := 15
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ans := containsNearbyDuplicate(nums, k)
	fmt.Println(ans)
}

func containsNearbyDuplicate(nums []int, k int) bool {
	n := len(nums)
	m := make(map[int]int, k+1)
	for i := 0; i < k; i++ {
		if i == n {
			return false
		}
		_, b := m[nums[i]]
		if b {
			return true
		}
		m[nums[i]]++
	}

	for i := k; i < n; i++ {
		_, b := m[nums[i]]
		if b {
			return true
		}
		m[nums[i]]++
		delete(m, nums[i-k])
	}
	return false
}
