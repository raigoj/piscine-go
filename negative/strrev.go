package piscine

func StrRev(s string) string {
	var revs string
	for i := len(s) - 1; i >= 0; i-- {
		revs = revs + string(s[i])
	}
	return revs
}
