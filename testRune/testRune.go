package main

import "fmt"

func main() {
	//s := "中国"
	s := "abcfffffff"
	r := []rune(s)
	fmt.Println(r)
	a := s[:]
	fmt.Println(a)
	fmt.Println(len(s[:]))
}
