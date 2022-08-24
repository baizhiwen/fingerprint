package main

import (
	"fmt"
	"sort"
)

func main() {
	envelopes := [][]int{{46,89},{50,53},{52,68},{72,45},{77,81}}
	ans := maxEnvelopes(envelopes)
	fmt.Println(ans)
}



//dp+二分查找
//自己写排序

type Env [][]int

func (v Env) Len() int {
	return len(v)
}

func (v Env) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v Env) Less(i, j int) bool {
	a,b := v[i],v[j]
	return a[0]<b[0] || (a[0]==b[0] && a[1]>b[1])
}
func maxEnvelopes(a [][]int) int {
	envelopes := Env{}
	envelopes = a
	sort.Sort(envelopes)
	n := len(envelopes)
	dp :=make([]int,0)

	//sort.Slice(envelopes, func(i, j int) bool {
	//	a,b := envelopes[i],envelopes[j]
	//	return a[0]<b[0] || (a[0]==b[0] && a[1]>b[1])
	//})

	ansMax := 0
	for i:=0; i < n ;i++  {
		if j := sort.SearchInts(dp, envelopes[i][1]); j < len(dp) {
			dp[j] = envelopes[i][1]
		} else {
			ansMax++
			dp = append(dp, envelopes[i][1])
		}
	}
	return ansMax
}

//dp+二分查找
func maxEnvelopes2(envelopes [][]int) int {
	n := len(envelopes)
	dp :=make([]int,n+1)
	sort.Slice(envelopes, func(i, j int) bool {
		a,b := envelopes[i],envelopes[j]
		return a[0]<b[0] || (a[0]==b[0] && a[1]>b[1])
	})
	ansMax := 1
	dp[ansMax]=envelopes[0][1]
	for i:=1; i < n ;i++  {
		if dp[ansMax] < envelopes[i][1] {
			ansMax++
			dp[ansMax] = envelopes[i][1]
		}else {
			var l,r,pos =1,ansMax,0
			for l<=r {
				mid := (r+l)/2
				if dp[mid] < envelopes[i][1] {
					pos = mid
					l = mid+1
				}else {
					r = mid-1
				}
			}
			dp[pos+1] = envelopes[i][1]
		}
	}
	return ansMax
}

//dp 超时
func maxEnvelopes1(envelopes [][]int) int {
	n := len(envelopes)
	dp :=make([]int,n)
	sort.Slice(envelopes, func(i, j int) bool {
		a,b := envelopes[i],envelopes[j]
		return a[0]<b[0] || (a[0]==b[0] && a[1]>b[1])
	})
	ans := 1
	for i:=0; i<n;i++{
		dp[i]=1
		for j:=0;j<i ;j++  {
			if envelopes[j][1] < envelopes[i][1]{
				dp[i] = maxNum1(dp[i],dp[j]+1)
			}
		}
		ans = maxNum1(ans,dp[i])
	}
	return ans
}

func maxNum1(x,y int) int {
	if x >y {
		return x
	}
	return y
}