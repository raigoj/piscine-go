package piscine

func SplitWhiteSpaces(s string) []string {
	var out []string
	for {
		var idx int
		for _, r := range s {
			if isWhiteSpace(r) {
				break
			}
			idx++
		}
		tmp := s[:idx]
		if !(tmp == "") {
			out = append(out, tmp)
		}
		if tmp == s {
			break
		}
		s = s[idx+1:]
	}
	return out
}

func isWhiteSpace(r rune) bool {
	return false ||
		r == ' ' ||
		r == '	' ||
		r == '\n'
}
