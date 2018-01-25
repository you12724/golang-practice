package main

import "testing"

func TestMax1(t *testing.T) {
	for _, test := range []struct {
		input  []int
		result int
	}{
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{1}, 1},
		{[]int{}, 0},
		{[]int{-1}, -1},
		{[]int{-1, 0}, 0},
		{[]int{-1, 2, 3, 4, -5}, 4},
	} {
		output := max1(test.input...)
		if test.result != output {
			t.Errorf("output is %v, expected %v", output, test.result)
		}
	}
}

func TestMax2(t *testing.T) {
	for _, test := range []struct {
		input1 int
		input2 []int
		result int
	}{
		{1, []int{2, 3, 4, 5}, 5},
		{1, []int{}, 1},
		{-1, []int{}, -1},
		{-1, []int{0}, 0},
		{-1, []int{2, 3, 4, -5}, 4},
	} {
		output := max2(test.input1, test.input2...)
		if test.result != output {
			t.Errorf("output is %v, expected %v", output, test.result)
		}
	}
}

func TestMin1(t *testing.T) {
	for _, test := range []struct {
		input  []int
		result int
	}{
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{1}, 1},
		{[]int{}, 0},
		{[]int{-1}, -1},
		{[]int{-1, 0}, -1},
		{[]int{-1, 2, 3, 4, -5}, -5},
	} {
		output := min1(test.input...)
		if test.result != output {
			t.Errorf("output is %v, expected %v", output, test.result)
		}
	}
}

func TestMin2(t *testing.T) {
	for _, test := range []struct {
		input1 int
		input2 []int
		result int
	}{
		{1, []int{2, 3, 4, 5}, 1},
		{1, []int{}, 1},
		{-1, []int{}, -1},
		{-1, []int{0}, -1},
		{-1, []int{2, 3, 4, -5}, -5},
	} {
		output := min2(test.input1, test.input2...)
		if test.result != output {
			t.Errorf("output is %v, expected %v", output, test.result)
		}
	}
}
