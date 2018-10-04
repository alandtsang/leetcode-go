// 托普利茨矩阵
// 除左下角[M-1, 0] 和右上角[0, N-1] 元素外，其余左上到右下对角线元素都相同
// 遍历的话，就是当前元素[i, j] 和 右下方对角线元素[i+1, j+1] 相等
package istoeplitzmatrix

func isToeplitzMatrix(matrix [][]int) bool {
	for i := 0; i < len(matrix)-1; i++ {
		for j := 0; j < len(matrix[0])-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}
	return true
}
