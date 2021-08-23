package incrementor

import (
	"math"
	"testing"
)

func TestNewIncrementor(t *testing.T) {
	inc := NewIncrementor()
	var want int32 = 0
	ans := inc.GetNumber()
	if ans != want {
		t.Error("NewIncrementor: wrong initialization")
	}

	var i int32 = 0
	for ; i < math.MaxInt32; i++ {
		inc.IncrementNumber()
	}
	inc.IncrementNumber()
	ans = inc.GetNumber()
	if ans != want {
		t.Error("NewIncrementor: maximum value exceeded")
	}
}

func TestGetNumber(t *testing.T) {
	inc := NewIncrementor()
	var want int32 = 0
	ans := inc.GetNumber()
	if ans != want {
		t.Error("GetNumber failed")
	}
}

func TestIncrementNumber(t *testing.T) {
	inc := NewIncrementor()
	var want int32 = 100
	var i int32 = 0
	for ; i < want; i++ {
		inc.IncrementNumber()
	}
	ans := inc.GetNumber()
	if ans != want {
		t.Error("IncrementNumber: wrong increment", ans)
	}
}

func TestSetMaximumValue(t *testing.T) {
	inc := NewIncrementor()
	var (
		maximumValue int32 = 5
		want         int32 = 1
	)
	inc.SetMaximumValue(maximumValue)
	var i int32 = 0
	for ; i <= maximumValue+want; i++ {
		inc.IncrementNumber()
	}
	ans := inc.GetNumber()
	if ans != want {
		t.Error("SetMaximumValue: wrong change of maximum value")
	}

	inc.SetMaximumValue(100)
	for i = 0; i < 50; i++ {
		inc.IncrementNumber()
	}

	want = 0
	inc.SetMaximumValue(10)
	ans = inc.GetNumber()
	if ans != want {
		t.Error("SetMaximumValue: wrong change of value")
	}
	
	for i = 0; i < 10; i++ {
		inc.IncrementNumber()
	}
	inc.IncrementNumber()
	ans = inc.GetNumber()
	if ans != want {
		t.Error("SetMaximumValue: wrong change of maximum value")
	}

	inc.SetMaximumValue(-40)
	inc.IncrementNumber()
	ans = inc.GetNumber()
	if ans < 0 {
		t.Error("SetMaximumValue: maximum value is negative")
	}

}