package addbinary

func addBinary(a string, b string) string {
	var carry, total uint8
	var result string
	lenA := len(a)
	lenB := len(b)

	for lenA > 0 && lenB > 0 {
		total = a[lenA-1] - '0' + b[lenB-1] - '0' + carry // 注意减 '0'
		carry = total / 2
		total = total % 2
		result += string(total + '0')
		lenA--
		lenB--
	}

	for lenA > 0 { // 处理 a 剩余字符
		total = a[lenA-1] - '0' + carry
		carry = total / 2
		total = total % 2
		result += string(total + '0')
		lenA--
	}

	for lenB > 0 { // 处理 b 剩余字符
		total = b[lenB-1] - '0' + carry
		carry = total / 2
		total = total % 2
		result += string(total + '0')
		lenB--
	}

	if carry != 0 {
		result += string(carry + '0')
	}
	return Reverse(result) // 得到的 result 是倒序的，需翻转
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
