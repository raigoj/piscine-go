package piscine

func Index(s string, toFind string) int {
	// if len(s) < len(toFind) {
	// 	return -1
	// }
	str := []rune(s)
	match := []rune(toFind)
	for i := 0; i < len(str)-len(match); i++ {
		if equal(str[i:i+len(match)], match) {
			return i
		}
	}
	return -1
}

func equal(a, b []rune) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
