package piscine

func LastRune(s string) rune {
	return rune([]rune(s)[len(s)-1])
}
