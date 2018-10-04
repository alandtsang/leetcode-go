// 遍历数组，pos 记录需要填入的位置
// 找到非 0 值，将其填入 nums[pos]，pos++
// 遍历到尾部后，将 pos 后的元素清 0
package movezeroes

func moveZeroes(nums []int) {
	var i, pos int // pos 指向被填入位置
	firstZero := true

	for i = 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if firstZero {
				pos = i // 记录第 1 个 0 的位置
				firstZero = false
			}
		} else { // 移动非 0 值到 nums[pos]
			nums[pos] = nums[i]
			pos++
			//fmt.Println(nums)
		}
	}

	for i = pos; i < len(nums); i++ { // 将 pos 之后元素都清 0
		nums[i] = 0
	}
}
