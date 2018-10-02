// 初始化当前最大和与总最大和为数组第一个元素
// 遍历数组，如果当前最大和小于 0 则更新其值为下个元素
// 如果当前最大和大于 0 则与下个元素累和
// 对比当前最大和与总最大和，若大于则更新总共最大和

package maxsubarray

func maxSubArray(nums []int) int {
	var length int
	if length = len(nums); length == 0 {
		return 0
	}

	var curSum, totalSum int
	curSum = nums[0]
	totalSum = nums[0]
	for i := 1; i < len(nums); i++ {
		if curSum < 0 {
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}

		if curSum > totalSum {
			totalSum = curSum
		}
	}
	return totalSum
}
