// 从 0 ~ n 共 n+1 个数，求和 (0 + n) * (n + 1) / 2
// 然后减去数组中的元素，结果是缺失的
package missingnumber

func missingNumber(nums []int) int {
	var n, total int
	if n = len(nums); n == 0 {
		return 0
	}

	total = (0 + n) * (n + 1) / 2
	for _, v := range nums {
		total -= v
	}
	return total
}
