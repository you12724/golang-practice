// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package cycle

import (
	"testing"
)

func TestIsCycle(t *testing.T) {
	type CyclePtr *CyclePtr
	var cyclePtr1, cyclePtr2 CyclePtr
	cyclePtr1 = &cyclePtr1
	cyclePtr2 = &cyclePtr2

	type CycleSlice []CycleSlice
	var cycleSlice = make(CycleSlice, 1)
	cycleSlice[0] = cycleSlice

	for _, test := range []struct {
		x    interface{}
		want bool
	}{
		// basic types
		{1, false},
		{1, false}, // different values
		{1, false}, // different types
		// slices
		{[]string{"foo"}, false},
		{[]string{"foo"}, false},
		{[]string{}, false},
		// slice cycles
		{cycleSlice, true},
		// pointer cycles
		{cyclePtr1, true},
		{cyclePtr2, true},
		{cyclePtr1, true}, // they're deeply equal
	} {
		if IsCycle(test.x) != test.want {
			t.Errorf("IsCycle(%v) = %t",
				test.x, !test.want)
		}
	}
}
