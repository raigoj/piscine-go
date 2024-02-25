package piscine

func ToUpper(s string) string {
	ss := []rune(s)
	for i, v := range ss {
		if 'a' <= v && v <= 'z' {
			ss[i] = rune(v - 32)
		}
	}
	return string(ss)
}
