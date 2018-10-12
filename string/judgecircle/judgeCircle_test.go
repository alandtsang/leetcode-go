package judgecircle

import "testing"

func TestJudgeCircle(t *testing.T) {
	if judgeCircle("UD") != true {
		t.Error("Get false, Expect true")
	}

	if judgeCircle("LL") != false {
		t.Error("Get true, Expect false")
	}

	if judgeCircle("LDRRLRUULR") != false {
		t.Error("Get true, Expect false")
	}
}
