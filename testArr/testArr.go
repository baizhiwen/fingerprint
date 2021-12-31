package main

import (
	"fmt"
	"reflect"
	"sort"
	"unsafe"
)

func main() {

	for i, c := range []byte("世界 abc") {
		fmt.Println(i, c)
	}

	var a = []string{"1.0", "1.1", "2.1.1", "1.2.1", "1.15.1", "1.3.1"}

	sort.Strings(a)

	fmt.Println(a)

	//arr := make([]string,0)
	//len(arr)
	//cap(arr)
	//reflect.SliceHeader
}

var a = []float64{4, 2, 5, 7, 2, 1, 88, 1}

func SortFloat64FastV1(a []float64) {
	// 强制类型转换
	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]

	// 以int方式给float64排序
	sort.Ints(b)
}

func SortFloat64FastV2(a []float64) {
	// 通过 reflect.SliceHeader 更新切片头部信息实现转换
	var c []int
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	*cHdr = *aHdr

	// 以int方式给float64排序
	sort.Ints(c)
}
