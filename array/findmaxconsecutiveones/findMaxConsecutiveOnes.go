// 遍历 nums，计算 1 的个数
// 当遇到 0 时，比较当前 1 个数和最大 1 个数
package findmaxconsecutiveones

func findMaxConsecutiveOnes(nums []int) int {
	var result, count int

	for _, v := range nums {
		if v == 1 {
			count++
		} else {
			if count > result {
				result = count
			}
			count = 0
		}

	}

	if count > result {
		result = count
	}
	return result
}
