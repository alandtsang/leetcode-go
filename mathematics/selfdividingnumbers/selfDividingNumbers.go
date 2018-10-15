// 用变量 i 遍历 left 到 right，i 对每个 i 拆分出来的数字求余
// 注意处理余数为 0 情况
package selfdividingnumbers

func selfDividingNumbers(left int, right int) []int {
	var result []int

	for i := left; i <= right; i++ {
		if divide(i) {
			result = append(result, i)
		}
	}
	return result
}

func divide(num int) bool {
	var n, remainder int
	n = num

	for n != 0 {
		remainder = n % 10
		if remainder == 0 || num%remainder != 0 { // 余数为 0 或无法除开
			return false
		}
		n /= 10
	}
	return true
}
