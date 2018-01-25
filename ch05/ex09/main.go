// 後回し
package main

func main() {
}

func expand(s string, f func(string) string) string {
	str := f("foo")
	target := "$foo"
	var result string
	var continueCount int
	for i, ele := range s {
		if continueCount > 0 {
			continueCount--
			continue
		}
		if i+len(str) > len(s) {
			result += s[i:len(s)]
			break
		}

		if s[i:i+len(target)] == target {
			result += str
			continueCount += len(target) - 1
		} else {
			result += string(ele)
		}
		// fmt.Printf("%s, %s\n", s[i:i+len(target)], result)
	}
	return result
}
