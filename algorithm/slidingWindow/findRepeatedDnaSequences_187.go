package main

import "fmt"

func main() {
	s := "AAAAAAAAAAA"
	arr := findRepeatedDnaSequences(s)
	fmt.Println(arr)
}
func findRepeatedDnaSequences(s string) []string {

	m := make(map[string]int, 0)
	n := len(s)
	ans := make([]string, 0)
	for i := 0; i <= n-10; i++ {
		sub := s[i : i+10]
		m[sub] = m[sub] + 1
	}
	for k, v := range m {
		if v > 1 {
			ans = append(ans, k)
		}
	}

	return ans
}
