package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	var isNeg bool
	if s[0] == '-' {
		isNeg = true
		s = s[1:]
	}
	var total int
	for _, v := range s {
		if !('0' <= v && v <= '9') {
			total = 0
			break
		}
		n := int(v - '0')
		total *= 10
		total += n
	}
	if isNeg {
		total = -total
	}
	return total
}
