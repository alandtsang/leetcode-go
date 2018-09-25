package addstrings

func addStrings(num1 string, num2 string) string {
	var i, j int
	var carry, total uint8
	var result string

	for i, j = len(num1)-1, len(num2)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		total = num1[i] - '0' + num2[j] - '0' + carry
		carry = total / 10
		total = total % 10
		result += string(total + '0')
	}

	for i >= 0 { // 处理 num1 剩余字符
		total = num1[i] - '0' + carry
		carry = total / 10
		total = total % 10
		result += string(total + '0')
		i--
	}

	for j >= 0 { // 处理 num2 剩余字符
		total = num2[j] - '0' + carry
		carry = total / 10
		total = total % 10
		result += string(total + '0')
		j--
	}

	if carry != 0 {
		result += string(carry + '0')
	}
	return Reverse(result)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
