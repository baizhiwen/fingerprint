package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 2, 2, 2, 2, 2, 2, 0, -1, 4, 4}

	fmt.Println(majorityElementQ(nums1))
}

func partition(data []int, l int, h int) int {
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
func quickSort(data []int, l int, h int) {
	if l < h {
		mid := partition(data, l, h)
		quickSort(data, l, mid-1)
		quickSort(data, mid+1, h)
	}
}

func majorityElementQ(nums []int) bool {
	quickSort(nums, 0, len(nums)-1)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false

}
func majorityElement(nums []int) int {

	major := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			major = num
		}
		if major == num {
			count += 1
		} else {
			count -= 1
		}
	}

	return major

}

func majorityElement2(nums []int) int {

	map1 := make(map[int]int, 0)
	for _, v := range nums {
		map1[v] = map1[v] + 1
	}
	ma1 := 0
	i := 0
	for k, v := range map1 {
		if v > ma1 {
			ma1 = v
			i = k
		}
	}
	return i

}

func majorityElement3(nums []int) bool {

	map1 := make(map[int]int, 0)
	for _, v := range nums {
		map1[v] = v
	}
	if len(map1) < len(nums) {
		return true
	}
	return false

}

func majorityElement1(nums []int) int {
	max := 0
	min := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}

	}
	bucket1 := make([]int, max+1)
	bucket2 := make([]int, -min+1)

	for _, v := range nums {
		if v < 0 {
			bucket2[-v]++
		} else {
			bucket1[v]++
		}

	}
	ma1 := 0
	ma2 := 0
	i := 0
	j := 0
	for k, v := range bucket1 {
		if v > ma1 {
			ma1 = v
			i = k
		}
	}
	for k, v := range bucket2 {
		if v > ma2 {
			ma2 = v
			j = k
		}
	}
	if ma2 > ma1 {
		return -j
	}
	return i

}
func majorityElement11(nums []int) bool {
	max := 0
	min := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}

	}
	bucket1 := make([]int, max+1)
	bucket2 := make([]int, -min+1)

	for _, v := range nums {
		if v < 0 {
			bucket2[-v]++
		} else {
			bucket1[v]++
		}

	}
	for _, v := range bucket1 {
		if v > 1 {
			return true
		}
	}
	for _, v := range bucket2 {
		if v > 1 {
			return true
		}
	}
	return false

}
