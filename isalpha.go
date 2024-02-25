package piscine

func IsAlpha(s string) bool {
	for _, v := range s {
		if !(('A' <= v && v <= 'Z') || ('a' <= v && v <= 'z') || ('0' <= v && v <= '9')) {
			return false
		}
	}
	return true
}
