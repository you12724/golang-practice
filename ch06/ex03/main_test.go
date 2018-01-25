package main

import "testing"

func TestIntersectWith(t *testing.T) {
	set1 := new(IntSet)
	set1.AddAll(1, 2, 3, 4)

	set2 := new(IntSet)
	set2.AddAll(3, 4, 5, 6)

	set1.IntersectWith(set2)

	if set1.Len() != 2 {
		t.Errorf("length is %d, expected 2", set1.Len())
	}

	if !(set1.Has(3) && set1.Has(4)) {
		t.Errorf("content is %v, expected 3, 4", set1)
	}
}

func TestDifferenceWith(t *testing.T) {
	set1 := new(IntSet)
	set1.AddAll(1, 2, 3, 4)

	set2 := new(IntSet)
	set2.AddAll(3, 4, 5, 6)

	set1.DifferenceWith(set2)

	if set1.Len() != 2 {
		t.Errorf("length is %d, expected 2", set1.Len())
	}

	if !(set1.Has(1) && set1.Has(2)) {
		t.Errorf("content is %v, expected 1, 2", set1)
	}
}

func TestSymmetricDifference(t *testing.T) {
	set1 := new(IntSet)
	set1.AddAll(1, 2, 3, 4)

	set2 := new(IntSet)
	set2.AddAll(3, 4, 5, 6)

	set1.SymmetricDifference(set2)

	if set1.Len() != 4 {
		t.Errorf("length is %d, expected 4", set1.Len())
	}

	if !(set1.Has(1) && set1.Has(2) && set1.Has(5) && set1.Has(6)) {
		t.Errorf("content is %v, expected 1, 2, 5, 6", set1)
	}
}
