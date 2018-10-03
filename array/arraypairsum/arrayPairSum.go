// 2n 数组两两组合，取每组最小值累和
// 当两个数比较接近的时取最小值才相对较大，否则会较小
// 比如 min(1, 4) 和 min(2, 3) 的值就分别是 1 和 2
// 但是如果 min(1, 2) 和 min(3, 4) 则值分别是 1 和 3
// 所以数组有序时，取偶数下标累和的结果最大
package arraypairsum

import (
	"sort"
)

func arrayPairSum(nums []int) int {
	var result int
	sort.Ints(nums)
	for i := 0; i < len(nums); i = i + 2 {
		result += nums[i]
	}
	return result
}
