package reverse

import "strconv"

func reverse(x int) int {
	var minusSign bool

	if x < 0 {
		minusSign = true
		x = -x
	}

	str := strconv.Itoa(x)
	str2 := reverseString(str)
	ret, err := strconv.Atoi(str2)
	if err != nil {
		return 0
	}

	if minusSign {
		ret = -ret
	}

	if ret > 2147483647 || ret < -2147483648 {
		return 0
	}
	return ret
}

func reverseString(s string) string {
	str := []rune(s)

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return string(str)
}

func reverse2(x int) int {
	if x == 0 {
		return x
	}

	var ret int
	for x != 0 {
		ret = ret*10 + x%10
		x = x / 10
	}

	if ret > 2147483647 || ret < -2147483648 {
		return 0
	}
	return ret
}
