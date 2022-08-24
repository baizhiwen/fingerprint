package main

func main() {
	//nums := []int{0,0,1,1,1,2,2,2,3,4,6}
	//l := removeElement(nums,2)
	//for i:= 0;i<l ; i++ {
	//	println(nums[i])
	//}
}

//27
func removeElement(nums []int, val int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	count := 0
	j := 0
	for i := 0; i < length; i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
			count++
		}
	}
	return count
}