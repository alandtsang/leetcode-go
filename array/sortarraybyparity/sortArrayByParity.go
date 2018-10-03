// 偶数放在数组的左侧，奇数防在右侧
// 可以从数据的左右两边向中间遍历
// 左侧找到奇数，右边找到偶数，进行交换
package sortarraybyparity

func sortArrayByParity(A []int) []int {
	for i, j := 0, len(A)-1; i < j; {
		if A[i]%2 == 0 { // 寻找左侧奇数
			i++
			continue
		}
		if A[j]%2 != 0 { // 寻找右侧偶数
			j--
			continue
		}
		A[i], A[j] = A[j], A[i] // 交换
	}
	return A
}
