package SumRange_304

func main() {

}

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	num := NumMatrix{}
	sum := make([][]int,len(matrix))
	for i:= 0; i< len(matrix);i++  {
		sum[i] = make([]int,len(matrix[i])+1)
		for j:=0; j< len(matrix[i]);j++  {
			sum[i][j+1]=sum[i][j]+matrix[i][j]
		}
	}
	num.matrix = sum
	return num
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		sum += this.matrix[i][col2+1] - this.matrix[i][col1]
	}
	return sum
}