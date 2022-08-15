package main

import (
	"fmt"
	"unsafe"
)

func main() {
	x, y := 6, 8

	//x,y = y,x

	change1(&x, &y)
	fmt.Println(x, y)

	//fun3()
}

func change1(x, y *int) {
	*x, *y = *y, *x
}

func fun3() {
	s := "asdfasdffff"
	b := string2bytes(s)
	fmt.Println(b)

	ss := bytes2string(b)
	fmt.Println(ss)
}
func string2bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
func bytes2string(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func fun2() {
	mp := make(map[string]int)
	mp["qcrao"] = 100
	mp["stefno"] = 18
	var l = (**int)(unsafe.Pointer(&mp))
	**l = 3
	var count = **l
	fmt.Println(count, len(mp)) // 2 2
	for _, v := range mp {
		println(v)
	}

}

func fun1() {
	s := make([]int, 9, 20)
	fmt.Println(len(s))
	var l = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	*l = 10
	var Len = *l
	fmt.Println(Len, len(s)) // 9 9
}
