// 杨辉三角
// 每行除两边元素为 1 外，每个元素 nums[i][j] = nums[i-1][j-1]+nums[i-1][j]

package pascalstriangle

func generate(numRows int) [][]int {
	var result [][]int
	for i := 1; i <= numRows; i++ {
		var row []int
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				row = append(row, 1)
			} else {
				row = append(row, result[i-2][j-1]+result[i-2][j])
			}
		}
		result = append(result, row)
	}
	return result
}
