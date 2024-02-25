package piscine

func IsUpper(s string) bool {
	for _, v := range s {
		if !('A' <= v && v <= 'Z') {
			return false
		}
	}
	return true
}
