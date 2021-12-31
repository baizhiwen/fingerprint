package main

import "fmt"

func main() {
	rev := isValid("{}{[]}()")
	fmt.Println(rev)
}

func isValid(s string) bool {
	l := len(s)

	if s == "" || l%2 != 0 {
		return false
	}

	valMap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack1 := []byte{}
	for i := 0; i < l; i++ {
		if valMap[s[i]] > 0 {
			if len(stack1) == 0 || stack1[len(stack1)-1] != valMap[s[i]] {
				return false
			}
			stack1 = stack1[:len(stack1)-1]
		} else {
			stack1 = append(stack1, s[i])
		}
	}
	return len(stack1) == 0
}
