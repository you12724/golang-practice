package main

import "testing"

func TestMain(t *testing.T) {
	set := new(IntSet)

	set.Add(10)
	if !set.Has(10) {
		t.Error("cant add")
	}

	set.Remove(10)
	if set.Has(10) {
		t.Error("cant remove")
	}

	set.Add(11)
	set.Clear()
	if set.Has(11) {
		t.Error("cant clear")
	}

	set.Add(1)
	set.Add(2)
	set.Add(3)
	if set.Len() != 3 {
		t.Errorf("length is %v, expected is 3", set.Len())
	}

	newset := set.Copy()
	newset.Add(4)
	if newset.Len() != 4 && set.Len() != 3 {
		t.Errorf("oldset is %v, newset is #v", set, newset)
	}
}
