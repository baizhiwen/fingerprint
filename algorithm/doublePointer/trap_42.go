package main

import "fmt"

func main() {
	height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	ans := trap(height)
	fmt.Println(ans)
}

func trap(height []int) int {
	l,r := 0,len(height)-1
	lmax,rmax := 0,0
	ans := 0
	for l<r  {
		lmax = max(lmax,height[l])
		rmax = max(rmax,height[r])
		if height[l] < height[r]{
			ans = ans + lmax-height[l]
			l++
		}else {
			ans = ans + rmax-height[r]
			r--
		}
	}
	return ans
}

func max(x,y int) int {
	if x > y {
		return x
	}
	return y
}