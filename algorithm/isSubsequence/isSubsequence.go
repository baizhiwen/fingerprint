package main

import "fmt"

func main() {
	s := "acfh"
	t := "abcdeeefggg"
	isS := isSubsequence(s, t)
	fmt.Println(isS)
}

//392判断子序列
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	sArr := []byte(s)
	tArr := []byte(t)
	j := 0
	for i := 0; i < len(tArr); i++ {
		if tArr[i] == sArr[j] {
			j++
			if j == len(sArr) {
				break
			}
		}
	}
	if j == len(sArr) {
		return true
	} else {
		return false
	}
}
