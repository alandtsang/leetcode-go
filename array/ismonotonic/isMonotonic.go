package ismonotonic

func isMonotonic(A []int) bool {
	var i, j int
	var increase, decrease bool

	increase = true
	decrease = true
	for i = 0; i < len(A)-1; i++ {
		j = i + 1
		if A[i] > A[j] {
			increase = false
			break
		}
	}

	for i = 0; i < len(A)-1; i++ {
		j = i + 1
		if A[i] < A[j] {
			decrease = false
			break
		}
	}
	return increase || decrease
}
