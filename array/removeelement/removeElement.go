package removeelement

func removeElement(nums []int, val int) int {
	var count int // 重复 val 个数
	for i := 0; i < len(nums); {
		if nums[i] == val {
			count++
			if i == len(nums)-1 { // 最后一个元素也等于 val
				nums = append(nums[:i-count+1])
			}
		} else {
			if count != 0 {
				nums = append(nums[:i-count], nums[i:]...)
				i = i - count // 更新 i
				count = 0
				continue
			}
		}
		i++
	}
	return len(nums)
}
