package number

import (
	"luago/number"
	"testing"
)

func assert(t *testing.T, s string, f func() bool){
	if !f() {
		t.Error(s)
	}
}

func TestIFloorDiv(t *testing.T) {
	if number.IFloorDiv(11, 3) != 3 {
		t.Errorf("IFloorDiv error" )
	}
	if number.IFloorDiv(9, 3) != 3 {
		t.Errorf("IFloorDiv error" )
	}
	if number.IFloorDiv(-9, 3) != -3 {
		t.Errorf("IFloorDiv error" )
	}
	if number.IFloorDiv(9, -3) != -3 {
		t.Errorf("IFloorDiv error" )
	}
	if number.IFloorDiv(10, -3) != -4 {
		t.Errorf("IFloorDiv error" )
	}
}

func TestFFloorDiv(t *testing.T) {
	if number.FFloorDiv(1.2, 0.3) != 4 {
		t.Errorf("FFloorDiv error")
	}
	if number.FFloorDiv(1.3, 0.3) != 4 {
		t.Errorf("FFloorDiv error")
	}
	if number.FFloorDiv(-1.3, 0.3) != -5 {
		t.Errorf("FFloorDiv error")
	}
}

func TestIMod(t *testing.T) {
	if number.IMod(10 ,3) != 1 {
		t.Error("error imod")
	}
	if number.IMod(-10 ,3) != 2 {
		t.Error("error imod")
	}
	if number.IMod(10 ,-3) != -2 {
		t.Error("error imod")
	}
	if number.IMod(-10 ,-3) != -1 {
		t.Error("error imod")
	}
}