package main

import "fmt"

func main() {
	strArr := []string{"aaaa", "aaab", "c"}
	rev := longestCommonPrefix(strArr)
	fmt.Println(rev)
}

func longestCommonPrefix(strArr []string) string {
	l := len(strArr)
	if l == 0 {
		return ""
	}
	p := strArr[0]
	for i := 1; i < l; i++ {
		p = lcp(p, strArr[i])
		if len(p) == 0 {
			break
		}
	}
	return p
}

func lcp(str1, str2 string) string {
	l := 0
	if len(str1) < len(str2) {
		l = len(str1)
	} else {
		l = len(str2)
	}
	index := 0
	for index < l && str1[index] == str2[index] {
		index++
	}

	return str1[:index]
}
