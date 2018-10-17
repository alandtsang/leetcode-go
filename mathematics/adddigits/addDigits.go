// 如果 num 不是个位数就进行计算，
// 计算过程拆解数字相加，将得到的结果进行上面的判断。
package adddigits

func addDigits(num int) int {
	for num >= 10 {
		num = split(num)
	}
	return num
}

func split(num int) int {
	var result int
	for num > 0 {
		result += num % 10
		num /= 10
	}
	return result
}
