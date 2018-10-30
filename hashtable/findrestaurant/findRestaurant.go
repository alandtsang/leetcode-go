// 使用 map【string]int 记录 list1 字符串及其下标，
// 遍历 list2，取在 map 中的 str 的下标，进行累和判断，
// 如果比 sum 小，清空 same，重新放入，
// 如果等于 sum，直接放入。
package findrestaurant

func findRestaurant(list1 []string, list2 []string) []string {
	var same []string
	var list1Map map[string]int

	list1Map = make(map[string]int)
	for i, v := range list1 {
		list1Map[v] = i
	}

	var sum = 2001
	for i, v := range list2 {
		if value, ok := list1Map[v]; ok {
			if i+value < sum {
				sum = i + value
				same = same[:0]
				same = append(same, v)
			} else if i+value == sum {
				same = append(same, v)
			}
		}
	}
	return same
}
