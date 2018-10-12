// 遍历操作字符串，统计 UD 和 LR 的出现次数
// U, up++
// D, up--
// R, right++
// L, right--
// 所有操作后 up 和 right 都为 0 则返回 true，否则 false
package judgecircle

func judgeCircle(moves string) bool {
	var up, right int
	for _, v := range moves {
		switch v {
		case 85: // 'U'
			up++
		case 68: // 'D'
			up--
		case 82: // 'R'
			right++
		case 76: // 'L'
			right--
		default:
			return false
		}
	}
	return up == 0 && right == 0
}
