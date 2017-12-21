package main

import "testing"

func TestConfirmConst(t *testing.T) {
	if KB != 1000 {
		t.Fail()
	}
	if MB != 1000000 {
		t.Fail()
	}
	if GB != 1000000000 {
		t.Fail()
	}
	if TB != 1000000000000 {
		t.Fail()
	}
	if PB != 1000000000000000 {
		t.Fail()
	}
	if EB != 1000000000000000000 {
		t.Fail()
	}
	if ZB != 1000000000000000000000 {
		t.Fail()
	}
	if YB != 1000000000000000000000000 {
		t.Fail()
	}

}
