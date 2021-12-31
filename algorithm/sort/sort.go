package main

import (
	"fmt"
)

func main() {

	nums1 := []int{1, 2, 3, 2, 2, 10, 2, 2, 2, 2, 0, -1, 4, 4}
	//quickSort(nums1,0, len(nums1)-1)
	//InsertSort(nums1)
	//selectSort(nums1)
	//BubbleSort(nums1)
	//SellSort(nums1)
	//MergeSort1(nums1,0, len(nums1))
	MergeSort3(nums1, len(nums1))
	fmt.Println(nums1)
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

//2。插入排序
func InsertSort(data []int) {
	for i := 1; i < len(data); i++ {
		temp := data[i]
		j := i
		for ; j > 0 && temp < data[j-1]; j-- {
			data[j] = data[j-1]
		}
		data[j] = temp
	}
}

//3。选择排序
func selectSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				temp := data[j]
				data[j] = data[i]
				data[i] = temp
			}
		}
	}
}

//4。冒泡排序
func BubbleSort(data []int) {
	l := len(data)
	for i := l - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if data[j+1] < data[j] {
				//temp :=data[j+1]
				//data[j+1] = data[j]
				//data[j] = temp
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

//5。希尔排序
//增量序列折半的插入排序
func SellSort(data []int) {
	l := len(data)
	for step := l / 2; step >= 1; step = step / 2 {
		for i := step; i < l; i = i + step {
			temp := data[i]
			j := i - step
			for ; j >= 0 && temp < data[j]; j = j - step {
				data[j+step] = data[j]
			}
			data[j+step] = temp
		}
	}
}

//6.1。归并排序----自顶向下
func MergeSort1(data []int, l, r int) {
	if r-l > 1 {
		mid := l + (r-l+1)/2
		MergeSort1(data, l, mid)
		MergeSort1(data, mid, r)
		merge1(data, l, mid, r)
	}
}

func merge1(data []int, l, mid, r int) {
	lSize := mid - l
	rSize := r - mid
	newSize := lSize + rSize
	re := make([]int, 0, newSize)

	i, j := 0, 0
	for i < lSize && j < rSize {
		v1 := data[l+i]
		v2 := data[mid+j]
		if v1 < v2 {
			re = append(re, v1)
			i++
		} else if v1 > v2 {
			re = append(re, v2)
			j++
		} else {
			re = append(re, v1)
			i++
			re = append(re, v2)
			j++
		}
	}
	re = append(re, data[l+i:mid]...)
	re = append(re, data[mid+j:r]...)

	for k := 0; k < newSize; k++ {
		data[l+k] = re[k]
	}
}

//6.2。归并排序----自底向上
func MergeSort2(data []int, l, r int) {
	s := 1
	for r-l > s {
		for i := l; i < r; i += s << 1 {
			var ll = i
			var mid = ll + s
			var rr = ll + (s << 1)
			if mid > r {
				return
			}
			if rr > r {
				rr = r
			}
			merge1(data, ll, mid, rr)
		}
		s <<= 1
	}
}

//6.3。归并排序----自底向上优化版本
func InsertSortForMerge(list []int) {
	n := len(list)
	// 进行 N-1 轮迭代
	for i := 1; i <= n-1; i++ {
		deal := list[i] // 待排序的数
		j := i - 1      // 待排序的数左边的第一个数的位置

		// 如果第一次比较，比左边的已排好序的第一个数小，那么进入处理
		if deal < list[j] {
			// 一直往左边找，比待排序大的数都往后挪，腾空位给待排序插入
			for ; j >= 0 && deal < list[j]; j-- {
				list[j+1] = list[j] // 某数后移，给待排序留空位
			}
			list[j+1] = deal // 结束了，待排序的数插入空位
		}
	}
}

// 自底向上归并排序优化版本
func MergeSort3(array []int, n int) {
	// 按照三个元素为一组进行小数组排序，使用直接插入排序
	blockSize := 3
	a, b := 0, blockSize
	for b <= n {
		InsertSort(array[a:b])
		a = b
		b += blockSize
	}
	InsertSortForMerge(array[a:n])

	// 将这些小数组进行归并
	for blockSize < n {
		a, b = 0, 2*blockSize
		for b <= n {
			merge3(array, a, a+blockSize, b)
			a = b
			b += 2 * blockSize
		}
		if m := a + blockSize; m < n {
			merge3(array, a, m, n)
		}
		blockSize *= 2
	}
}

// 原地归并操作
func merge3(array []int, begin, mid, end int) {
	// 三个下标，将数组 array[begin,mid] 和 array[mid,end-1]进行原地归并
	i, j, k := begin, mid, end-1 // 因为数组下标从0开始，所以 k = end-1

	for j-i > 0 && k-j >= 0 {
		step := 0
		// 从 i 向右移动，找到第一个 array[i]>array[j]的索引
		for j-i > 0 && array[i] <= array[j] {
			i++
		}

		// 从 j 向右移动，找到第一个 array[j]>array[i]的索引
		for k-j >= 0 && array[j] <= array[i] {
			j++
			step++
		}

		// 进行手摇翻转，将 array[i,mid] 和 [mid,j-1] 进行位置互换
		// mid 是从 j 开始向右出发的，所以 mid = j-step
		rotation(array, i, j-step, j-1)
		i = i + step
	}
}

// 手摇算法，将 array[l,l+1,l+2,...,mid-2,mid-1,mid,mid+1,mid+2,...,r-2,r-1,r] 从mid开始两边交换位置
// 1.先逆序前部分：array[mid-1,mid-2,...,l+2,l+1,l]
// 2.后逆序后部分：array[r,r-1,r-2,...,mid+2,mid+1,mid]
// 3.上两步完成后：array[mid-1,mid-2,...,l+2,l+1,l,r,r-1,r-2,...,mid+2,mid+1,mid]
// 4.整体逆序： array[mid,mid+1,mid+2,...,r-2,r-1,r,l,l+1,l+2,...,mid-2,mid-1]
func rotation(array []int, l, mid, r int) {
	reverse(array, l, mid-1)
	reverse(array, mid, r)
	reverse(array, l, r)
}

func reverse(array []int, l, r int) {
	for l < r {
		// 左右互相交换
		array[l], array[r] = array[r], array[l]
		l++
		r--
	}
}
