package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	s := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	ans := minSubArrayLen3(s, nums)
	fmt.Println(ans)
}

//暴力，超时
func minSubArrayLen(s int, nums []int) int {
	ans := math.MaxInt32
	n := len(nums)
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum >= s {
				ans = min(ans, j-i+1)
				break
			}
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

//二分法
func minSubArrayLen2(s int, nums []int) int {
	ans := math.MaxInt32
	n := len(nums)
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	for i := 1; i <= n; i++ {
		t := s + sum[i-1]
		b := sort.SearchInts(sum, t)
		//if b < 0 {
		//	b = -b - 1
		//}
		if b <= n {
			ans = min(ans, b-i+1)
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

//滑动窗口
func minSubArrayLen3(s int, nums []int) int {
	ans := math.MaxInt32
	n := len(nums)
	sum := 0
	start, end := 0, 0
	for end < n {
		sum = sum + nums[end]
		for s <= sum {
			sum = sum - nums[start]
			ans = min(ans, end-start+1)
			start++
		}
		end++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
