// 元素范围 [1, n] 遍历数组，将元素对应下标的取反，目的：
// 1. 对应位取反做标记，最后未取反的就是缺失的数字
// 2. 取反后可以通过绝对值得到该位置的初始值
package finddisappearednumbers

import (
	"math"
)

func findDisappearedNumbers(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	var result []int
	var index int
	for i := 0; i < len(nums); i++ {
		//fmt.Printf("%2d ", nums[i])
		index = int(math.Abs(float64(nums[i]))) - 1 // 获取值的数组下标
		if nums[index] > 0 {
			nums[index] = -nums[index] // 取反保持原值可以通过 Abs 还原
		}
		//fmt.Printf("%2v\n", nums)
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}
	return result
}
