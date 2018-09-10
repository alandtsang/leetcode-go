package plusone

func plusOne(digits []int) []int {
	var s []int
	var carry, n int
	length := len(digits) - 1

	for i := length; i >= 0; i-- {
		if i == length {
			n = digits[i] + 1
		} else {
			n = digits[i] + carry
		}
		if n >= 10 {
			carry = n / 10
			n = n % 10
		} else {
			carry = 0
		}
		s = append(s, n)
	}

	if carry > 0 {
		s = append(s, carry)
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
