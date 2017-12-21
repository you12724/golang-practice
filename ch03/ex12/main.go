package main

import (
	"reflect"
	"sort"
	"strings"
)

func main() {

}

func judgeAnagram(input1, input2 string) bool {
	input1Arr := strings.Split(input1, "")
	input2Arr := strings.Split(input2, "")
	sort.Strings(input1Arr)
	sort.Strings(input2Arr)

	return reflect.DeepEqual(input1Arr, input2Arr)
}
