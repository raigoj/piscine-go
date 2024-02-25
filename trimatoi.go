package piscine

func TrimAtoi(s string) int {
	var neg bool
	var numRead bool
	var sBuf string
	for _, v := range s {
		if !numRead && v == '-' {
			neg = true
		}
		if '0' <= v && v <= '9' {
			numRead = true
			sBuf = sBuf + string(v)
		}
	}
	res := Atoi(sBuf)
	if neg {
		res *= -1
	}
	return res
}
