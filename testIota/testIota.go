package main

import "fmt"

type Status uint32

const (
	StatusOpen = iota
	StatusClosed
	StatusUnknown
)

func main() {
	fmt.Println(StatusOpen)
	fmt.Println(StatusClosed)
	fmt.Println(StatusUnknown)

	a := 23
	b := 12.8
	sum := a + int(b)
	fmt.Println(sum)

	fmt.Println(1 | 2 | 16)

	ss := fmt.Sprintf("aaaa,%s", "bbbb")
	fmt.Println(ss)

}
