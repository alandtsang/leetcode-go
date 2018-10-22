// candies 是偶数长度，平分的话，妹妹最多获得 result = len(candies) / 2 种类水果，
// 遍历 candies 用 map 存储种类，
// 如果种类超过 result 则返回 result，否则返回 map 长度。
package distributecandies

func distributeCandies(candies []int) int {
	var result int
	var candyMap map[int]int

	candyMap = make(map[int]int)
	for _, v := range candies {
		candyMap[v]++
	}
	result = len(candies) / 2
	if len(candyMap) < result {
		return len(candyMap)
	}
	return result
}
