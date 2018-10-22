package numjewelsinstones

import "testing"

type Tables struct {
	J   string
	S   string
	num int
}

func TestNumJewelsInStones(t *testing.T) {
	var result int
	tables := []Tables{
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}

	for _, table := range tables {
		result = numJewelsInStones(table.J, table.S)
		if result != table.num {
			t.Errorf("J=%s, S=%s, get %d, expect %d", table.J, table.S, result, table.num)
		}
	}
}
