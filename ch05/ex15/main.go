package main

func main() {

}

func max1(vals ...int) int {
	var result int
	for i, val := range vals {
		if i == 0 {
			result = val
			continue
		}

		if val > result {
			result = val
		}
	}

	return result
}

func min1(vals ...int) int {
	var result int
	for i, val := range vals {
		if i == 0 {
			result = val
			continue
		}

		if val < result {
			result = val
		}
	}

	return result
}

func max2(input int, vals ...int) int {
	var result int
	result = input
	for _, val := range vals {
		if val > result {
			result = val
		}
	}

	return result
}

func min2(input int, vals ...int) int {
	var result int
	result = input
	for _, val := range vals {
		if val < result {
			result = val
		}
	}

	return result
}
