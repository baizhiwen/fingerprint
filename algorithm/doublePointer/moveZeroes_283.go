package main

import "fmt"

func main() {
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int)  {
	l,r,n := 0,0, len(nums)
	for ;r<n;r++  {
		if nums[r] != 0 {
			nums[l],nums[r]=nums[r],nums[l]
			l++
		}
	}
}