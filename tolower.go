package piscine

func ToLower(s string) string {
	ss := []rune(s)
	for i, v := range ss {
		if 'A' <= v && v <= 'Z' {
			ss[i] = rune(v + 32)
		}
	}
	return string(ss)
}
