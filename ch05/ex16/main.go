package main

func main() {

}

func join(sep string, vals ...string) string {
	var result string
	for i, val := range vals {
		result += val
		if i != len(vals)-1 {
			result += sep
		}
	}
	return result
}
