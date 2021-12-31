package main

import "fmt"

func main() {
	rev := romanToInt("III")
	fmt.Println(rev)
}
func romanToInt(s string) int {
	romanMap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	l := len(s)
	rev := 0
	for i := range s {
		val := romanMap[s[i]]
		if i < l-1 && val < romanMap[s[i+1]] {
			rev = rev - romanMap[s[i]]
		} else {
			rev = rev + romanMap[s[i]]
		}
	}
	return rev
}
