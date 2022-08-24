package main

import "fmt"

func main() {
	piles := []int{30,11,23,4,20}
	h := 6
	ans := minEatingSpeed(piles,h)
	fmt.Println(ans)
}

func minEatingSpeed(piles []int, h int) int {
	max := 0
	for _,v := range piles{
		if v > max {
			max = v
		}
	}
	left, right := 1, max
	ans := max
	for left <= right {
		mid := (right-left)/2 + left
		time := getSpeed(piles,mid)
		if time <= h {
			ans = mid
			right = mid-1
		} else {
			left = mid + 1
		}
	}
	return ans
}

func getSpeed(piles []int,speed int) int {
	time := 0
	for i:=0;i< len(piles);i++ {
		s := (piles[i]+speed-1)/speed
		time+=s
	}
	return time
}