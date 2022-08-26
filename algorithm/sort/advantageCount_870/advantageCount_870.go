package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {
	//nums1 := []int{2,7,11,15}
	//nums2 := []int{1,10,4,11}
	nums1 := []int{2,0,4,1,2}
	nums2 := []int{1,3,0,0,2}

	ans := advantageCount(nums1,nums2)
	fmt.Println(ans)
}

//田忌赛马。这个效率高了些，比不过就直接匹配对面最大值
type B struct {
	x int
	y int
}
func advantageCount(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))
	numsB :=make([]B,0)
	for i:=0;i< len(nums2);i++ {
		b := B{nums2[i],i}
		numsB = append(numsB, b)
	}
	sort.Ints(nums1)
	sort.Slice(numsB, func(i, j int) bool {
		return numsB[i].x < numsB[j].x
	})
	l,r := 0,len(nums1)-1
	for i:=0;i< len(nums1);i++  {
		if nums1[i]>numsB[l].x {
			ans[numsB[l].y]=nums1[i]
			l++
		}else {
			ans[numsB[r].y]=nums1[i]
			r--
		}
	}
	return ans
}


//类似田忌赛马。排序：快排；map里面加链表，list链表的使用
//性能挺低
func advantageCount1(nums1 []int, nums2 []int) []int {
	ans := make([]int,0)
	less := make([]int,0)
	numsB :=make([]int,len(nums2))
	copy(numsB, nums2)
	sort.Ints(nums1)
	sort.Ints(numsB)
	j:=0
	bMap := make(map[int]*list.List,0)
	for i:=0;i< len(nums2);i++  {
		bMap[nums2[i]]=list.New()
	}
	for i:=0;i< len(nums1);i++ {
		if numsB[j] < nums1[i]{
			bMap[numsB[j]].PushBack(nums1[i])
			j++
		}else {
			less = append(less, nums1[i])
		}
	}
	j = 0
	for i:=0;i< len(nums2);i++ {
		if l := bMap[nums2[i]];l.Len()>0 {
			e := l.Front()
			l.Remove(e)
			ans = append(ans, (e.Value).(int))
		}else {
			ans = append(ans, less[j])
			j++
		}
	}
	return ans
}