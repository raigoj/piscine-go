package piscine

func Split(s, sep string) []string {
	var out []string
	for {
		i := index(s, sep)
		if i == -1 {
			break
		}
		if i > 0 {
			out = append(out, s[:i])
		}
		s = s[i+len(sep):]
	}
	if !(s == "") {
		out = append(out, s)
	}
	return out
}

func index(s string, toFind string) int {
	for i := 0; i <= len(s)-len(toFind); i++ {
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}
	return -1
}
