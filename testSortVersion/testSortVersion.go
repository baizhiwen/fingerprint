package main

import (
	"fmt"
	"gobase/utils/format"
	"sort"
	"strings"
)

var (
	version0 = 0
	version1 = 1
	version2 = 2
)

func StrTrimSpace(v1str, v2str string) (v1, v2 string) {
	v1 = strings.TrimSpace(v1str)
	v2 = strings.TrimSpace(v2str)
	return
}

func comparSlice(v1slice, v2slice []string) int {
	for index, _ := range v1slice {
		if format.StringToInt(v1slice[index]) > format.StringToInt(v2slice[index]) {
			return version1
		}
		if format.StringToInt(v1slice[index]) < format.StringToInt(v2slice[index]) {
			return version2
		}
		if len(v1slice)-1 == index {
			return version0
		}
	}
	return version0
}

func comparSlice1(v1slice, v2slice []string, flas int) int {
	for index, _ := range v1slice {
		//按照正常逻辑v1slice 长度小
		if format.StringToInt(v1slice[index]) > format.StringToInt(v2slice[index]) {
			if flas == 2 {
				return version2
			}
			return version1

		}
		if format.StringToInt(v1slice[index]) < format.StringToInt(v2slice[index]) {
			if flas == 2 {
				return version1
			}
			return version2
		}
		if len(v1slice)-1 == index {
			if flas == 2 {
				return version1
			} else if flas == 1 {
				return version2
			}
		}
	}
	return version0
}

func compareStrVer(v1, v2 string) (res int) {
	s1, s2 := StrTrimSpace(v1, v2)
	v1slice := strings.Split(s1, ".")
	v2slice := strings.Split(s2, ".")
	//长度不相等直接退出
	if len(v1slice) != len(v2slice) {
		if len(v1slice) > len(v2slice) {
			res = comparSlice1(v2slice, v1slice, 2)
			return res
		} else {
			res = comparSlice1(v1slice, v2slice, 1)
			return res
		}
	} else {
		res = comparSlice(v1slice, v2slice)
	}
	return res

}

/*
	简易版的版本号对比，要求必须版本号位数相同，否则对比不了，这里也是存在一个bug，这个版本已经解决
	1.1.1 --> 1.2.1 ok
	1.1.1 --> 1.2.12 ok
	1.1.1 -->  1.2   ok
	1.2 --> 3.2.2 ok
*/

func QuickSort(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	}
	splitdata := arr[0]           //第一个数据
	low := make([]string, 0, 0)   //比我小的数据
	hight := make([]string, 0, 0) //比我大的数据
	mid := make([]string, 0, 0)   //与我一样大的数据
	mid = append(mid, splitdata)  //加入一个
	for i := 1; i < len(arr); i++ {
		if compareStrVer(arr[i], splitdata) == 1 {
			low = append(low, arr[i])
		} else if compareStrVer(arr[i], splitdata) == 2 {
			hight = append(hight, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low, hight = QuickSort(low), QuickSort(hight)
	myarr := append(append(low, mid...), hight...)
	return myarr
}

type AA []string

func (v AA) Len() int {
	return len(v)
}

func (v AA) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v AA) Less(i, j int) bool {

	if compareStrVer(v[i], v[j]) == 2 {
		return true
	}

	return false
}

type bb struct {
}

func (b *bb) d() {

}

type CC interface {
	d()
}

func main() {
	var _ CC = (*bb)(nil)

	var a = []string{"1.0", "1.1", "2.1.1", "1.2.1", "1.15.1", "1.3.1"}

	b := QuickSort(a)
	fmt.Println(b)

	var aa = AA{}
	aa = []string{"1.0", "1.1", "2.1.1", "1.2.1", "1.15.1", "1.3.1"}

	sort.Sort(aa)

	fmt.Println(aa)

}
