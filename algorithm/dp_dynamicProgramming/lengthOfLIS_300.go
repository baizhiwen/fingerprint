package main

import "fmt"

func main() {
	nums := []int{10,9,2,5,3,7,101,18}
	ans := lengthOfLIS(nums)
	fmt.Println(ans)
}

//贪心 + 二分查找
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ansMax := 1
	dp := make([]int,n+1)
	dp[ansMax]=nums[0]
	for i:=1; i < n ;i++  {
		if dp[ansMax] < nums[i] {
			ansMax++
			dp[ansMax] = nums[i]
		}else {
			var l,r,pos =1,ansMax,0
			for l<=r {
				mid := (r+l)/2
				if dp[mid] < nums[i] {
					pos = mid
					l = mid+1
				}else {
					r = mid-1
				}
			}
			dp[pos+1] = nums[i]
		}
	}
	return ansMax
}

//dp
func lengthOfLIS2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ansMax := 1
	dp := make([]int,n)
	dp[0]=1
	for i:=1; i < n ;i++  {
		dp[i]=1
		for j:=0;j < i ; j++ {
			if nums[j] < nums[i] {
				dp[i] = maxNum(dp[i],dp[j]+1)
			}
		}
		ansMax = maxNum(ansMax,dp[i])
	}
	return ansMax
}

func maxNum(x,y int) int {
	if x > y {
		return x
	}
	return y
}