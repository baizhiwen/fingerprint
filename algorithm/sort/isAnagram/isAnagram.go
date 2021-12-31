package main

import (
	"fmt"
)

func main() {
	s1 := "kquci"
	s2 := "quick"

	fmt.Println(isAnagram(s1, s2))

}

func partition(data []byte, l int, h int) int {
	m := data[l]
	for l < h {
		for h > l && data[h] >= m {
			h--
		}
		if h > l {
			data[l] = data[h]
		}
		for l < h && data[l] <= m {
			l++
		}
		if l < h {
			data[h] = data[l]
		}
	}
	data[l] = m
	return l
}

//1。快排
func quickSort(data []byte, l int, h int) {
	if l < h {
		mid := partition(data, l, h)
		quickSort(data, l, mid-1)
		quickSort(data, mid+1, h)
	}
}

func isAnagram(s, t string) bool {
	s1 := []byte(s)
	s2 := []byte(t)
	quickSort(s1, 0, len(s1)-1)
	quickSort(s2, 0, len(s2)-1)
	return string(s1) == string(s2)
}
