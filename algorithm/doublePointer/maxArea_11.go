package main

import (
	"fmt"
)

func main() {
	h := []int{1,8,6,2,5,4,8,3,7}
	ans := maxArea(h)
	fmt.Println(ans)
}
//两端向中间靠，时间复杂度O(n)
func maxArea(height []int) int {
	maxA := 0
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	l,r := 0,len(height)-1
	for l<r  {
		area := (r-l)*min(height[r],height[l])
		if area > maxA {
			maxA = area
		}
		if height[l]<height[r]{
			l++
		}else {
			r--
		}
	}
	return maxA
}
//暴力超时
func maxArea1(height []int) int {
	maxA := 0
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	for i:=0; i< len(height)-1;i++  {
		for j:=i+1;j<len(height);j++  {
			area := (j-i)*min(height[i],height[j])
			if area > maxA {
				maxA = area
			}
		}
	}
	return maxA
}
