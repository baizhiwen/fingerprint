package SumRange_303

func main() {

}
type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	array := NumArray{}
	sum := make([]int, len(nums)+1)
	for i:=0;i< len(nums);i++  {
		sum[i+1] = sum[i]+nums[i]
	}
	array.nums = sum
	return array
}

func (this *NumArray) SumRange(left int, right int) int {
	nums := this.nums
	return nums[right+1]-nums[left]
}
//func Constructor(nums []int) NumArray {
//	array := NumArray{}
//	array.nums = nums
//	return array
//}
//
//func (this *NumArray) SumRange(left int, right int) int {
//	sum := 0
//	nums := this.nums
//	for i:= left; i<=right; i++ {
//		sum+=nums[i]
//	}
//	return sum
//}
