// 用 result 和 times 分别保存数组当前值以及出现的次数
// 遍历数组，如果 times 为 0 更新 result 和 times
// 如果下个元素和 result 相同，times++ 否则 times--
// 由于要找的元素出现的次数最多，所以最后 result 就行要找的值
package majorityelement

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var result, times int
	result = nums[0] // 数组第一个元素
	times = 1        // 元素出现的次数
	for i := 1; i < len(nums); i++ {
		if times == 0 {
			result = nums[i]
			times = 1
		} else if result == nums[i] {
			times++
		} else {
			times--
		}
	}
	return result
}
