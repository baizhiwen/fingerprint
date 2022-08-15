package main

import "fmt"

func main() {
	s := "abcabcabc"

	l := lengthOfLongestSubstring(s)
	fmt.Println(l)
}

//滑动窗口
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int,0)
	rk,ans :=-1,0
	n := len(s)
	for i:=0; i<n; i++ {
		if i != 0 {
			delete(m,s[i-1])
		}
		for rk+1 < n && m[s[rk+1]]==0 {
			m[s[rk+1]]++
			rk++
		}
		ans = max(ans,rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}