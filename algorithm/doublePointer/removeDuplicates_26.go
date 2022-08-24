package main

func main() {
	//nums := []int{0,0,1,1,1,2,2,2,3,4,6}
	//l := removeDuplicates(nums)
	//for i:= 0;i<l ; i++ {
	//	println(nums[i])
	//}


}

//26
func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	count := 1
	j := 0
	for i := 1; i < length; i++ {
		if nums[j] != nums[i] {
			count++
			j++
			nums[j] = nums[i]
		}
	}

	return count
}

