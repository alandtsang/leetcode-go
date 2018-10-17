package mycirculardeque

import "testing"

func TestMyCircularDeque(t *testing.T) {
	var circularDeque MyCircularDeque
	var value int

	circularDeque = Constructor(3)
	if circularDeque.InsertLast(1) != true { // 返回 true
		t.Errorf("Get false, Expect true")
	}
	if circularDeque.InsertLast(2) != true { // 返回 true
		t.Errorf("Get false, Expect true")
	}
	if circularDeque.InsertFront(3) != true { // 返回 true
		t.Errorf("Get false, Expect true")
	}
	if circularDeque.InsertFront(4) != false { // 已经满了，返回 false
		t.Errorf("Get true, Expect false")
	}
	if value = circularDeque.GetRear(); value != 2 { // 返回 2
		t.Errorf("Get %d, Expect 4", value)
	}
	if circularDeque.IsFull() != true { // 返回 true
		t.Errorf("Get false, Expect true")
	}
	if circularDeque.DeleteLast() != true { // 返回 true
		t.Errorf("Get false, Expect true")
	}
	if circularDeque.InsertFront(4) != true { // 返回 true
		t.Errorf("Get false, Expect true")
	}
	if value = circularDeque.GetFront(); value != 4 { // 返回 4
		t.Errorf("Get %d, Expect 4", value)
	}
}
