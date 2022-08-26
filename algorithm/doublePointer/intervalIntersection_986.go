package main

import "fmt"

func main() {
	firstList := [][]int{{0,2},{5,10},{13,23},{24,25}}
	secondList := [][]int{{1,5},{8,12},{15,24},{25,26}}
	ans := intervalIntersection(firstList,secondList)
	fmt.Println(ans)
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	ans := make([][]int,0)
	flen,slen := len(firstList), len(secondList)
	i,j := 0,0
	for i < flen&&j<slen {
		if firstList[i][1] < secondList[j][0] {
			i++
			continue
		}
		if firstList[i][0] > secondList[j][1] {
			j++
			continue
		}
		l := max1(firstList[i][0],secondList[j][0])
		r := min1(firstList[i][1],secondList[j][1])
		ans = append(ans, []int{l,r})
		if firstList[i][1] == secondList[j][1] {
			i++
			j++
		}else if firstList[i][1] < secondList[j][1] {
			i++
		}else {
			j++
		}
	}
	return ans
}

func max1(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func min1(x,y int) int {
	if x < y {
		return x
	}
	return y
}