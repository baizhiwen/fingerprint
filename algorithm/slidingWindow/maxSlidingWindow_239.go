package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	k := 3
	nums := []int{1,3,-1,-3,5,6,7}
	arr := maxSlidingWindow2(nums,k)
	fmt.Println(arr)
}

var a []int
type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}

func maxSlidingWindow2(nums []int, k int) []int {
	q := []int{}
	for i := 0;i<k;i++{
		q = push(q,i,nums)
	}
	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0]= nums[q[0]]
	for i := k; i < n; i++ {
		q = push(q,i,nums)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}

	return ans
}

func push(q []int,i int,nums []int) []int {
	for len(q) >0 && nums[q[len(q)-1]]< nums[i] {
		q = q[:len(q)-1]
	}
	q = append(q, i)
	return q
}
