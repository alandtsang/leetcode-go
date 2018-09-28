package minstack

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	var val int

	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	if val = obj.GetMin(); val != -3 {
		t.Errorf("Get %d, Expect -3", val)
	}
	obj.Pop()
	if val = obj.Top(); val != 0 {
		t.Errorf("Get %d, Expect 0", val)
	}
	if val = obj.GetMin(); val != -2 {
		t.Errorf("Get %d, Expect -2", val)
	}
}
