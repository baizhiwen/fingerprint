package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	h := &hp{make([]int, 0)}
	//for i:=0;i<3;i++ {
	//	h.IntSlice = append(h.IntSlice, list[i])
	//}

	heap.Init(h)
	for i := 0; i < len(list); i++ {
		//h.IntSlice = append(h.IntSlice, v)
		heap.Push(h, list[i])
	}
	// 将堆元素移除
	for i := 0; i < len(list); i++ {
		v := heap.Pop(h)
		fmt.Println(v)
	}

	// 打印排序后的值
	//fmt.Println(list)
}
