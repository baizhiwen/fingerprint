package main

import (
	"fmt"
	"math"
)

func main() {
	rev := reverse(9463847412)
	fmt.Println(rev)
	fmt.Println(math.MinInt32)
	fmt.Println(math.MaxInt32)
}
func reverse(x int) (rev int) {
	for x != 0 {
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
		if rev < math.MinInt32 || rev > math.MaxInt32 {
			return 0
		}
	}
	return
}
