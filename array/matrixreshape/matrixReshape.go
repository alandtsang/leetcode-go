// 矩阵转换可以先将原矩阵转为 1维矩阵
// 然后根据要求将 1维矩阵转为行和列
// 2维转1维 M[column*i+j] = M[i][j]
package matrixreshape

func matrixReshape(nums [][]int, r int, c int) [][]int {
	var result [][]int
	var row, column int
	row = len(nums)
	column = len(nums[0])

	if row*column != r*c { // 转换前后矩阵元素个数不等
		return nums
	}

	total := make([]int, r*c) // 转为1维数组
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			total[column*i+j] = nums[i][j]
		}
	}

	for i := 0; i < r; i++ { // 转为 [r, c] 矩阵
		var x []int
		for j := 0; j < c; j++ {
			x = append(x, total[c*i+j])
		}
		result = append(result, x)
	}
	return result
}
