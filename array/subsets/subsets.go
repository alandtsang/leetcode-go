// 全排列和部分排列问题都可以用回溯算法
// 注意：
// 1. 二维数组的修改需要传地址
// 2. 缓存的 temp 部分要用拷贝，否则会被修改
package subsets

import (
	"sort"
)

func subsets(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	backtrack(&result, []int{}, nums, 0)
	return result
}

func backtrack(result *[][]int, t, nums []int, start int) {
	temp := make([]int, len(t))
	copy(temp, t) // 拷贝到 temp，使用 t 的话会被修改
	*result = append(*result, temp)
	//fmt.Printf("-> append temp=%v, result=%v, start=%d\n", temp, *result, start)
	for i := start; i < len(nums); i++ {
		//fmt.Printf("### i=%d, temp=%v ###\n", i, temp)
		temp = append(temp, nums[i])
		//fmt.Println("  add, temp=", temp)
		backtrack(result, temp, nums, i+1)
		temp = append(temp[:len(temp)-1])
		//fmt.Println("  pop, temp=", temp)
	}
	//fmt.Printf("<- result=%v\n", *result)
}
