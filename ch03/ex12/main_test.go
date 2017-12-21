package main

import "testing"

func TestJudgeAnagram(t *testing.T) {
	input1 := "test"
	input2 := "ttes"
	if !judgeAnagram(input1, input2) {
		t.Fatal("fail...")
	}
}
