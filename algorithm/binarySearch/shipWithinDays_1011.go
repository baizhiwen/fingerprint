package main

import "fmt"

func main() {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	days := 5
	ans := shipWithinDays(weights, days)
	fmt.Println(ans)
}

func shipWithinDays(weights []int, days int) int {
	l,r := 0,0
	for i:=0;i< len(weights); i++ {
		if l < weights[i] {
			l = weights[i]
		}
		r = r+weights[i]
	}
	for l < r {
		mid := (r-l)/2+l
		need := 1
		cur := 0
		for i:=0;i< len(weights); i++{
			if cur+weights[i] > mid {
				need++
				cur=0
			}
			cur = cur+weights[i]
		}
		if need <= days {
			r = mid
		}else {
			l = mid+1
		}
	}
	return l
}