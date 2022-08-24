package main

import (
	"fmt"
	"math"
)

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	ans := minWindow(s, t)
	fmt.Println(ans)
}
func minWindow(s string, t string) string {
	ori, cnt := make(map[byte]int, 0), make(map[byte]int, 0)
	maxlen := math.MaxInt32
	slen := len(s)
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}
	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	ansl, ansr := -1, -1
	for l, r := 0, 0; r < slen; r++ {
		if ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < maxlen {
				maxlen = r - l + 1
				ansl, ansr = l, l+maxlen
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]]--
			}
			l++
		}
	}
	if ansl == -1 {
		return ""
	}
	return s[ansl:ansr]
}

func minWindow2(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < len {
				len = r - l + 1
				ansL, ansR = l, l+len
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}
