package piscine

func Capitalize(s string) string {
	var res string
	var upper bool = true
	for _, v := range s {
		switch {
		case ('a' <= v && v <= 'z'):
			if upper {
				v = v - 32
				upper = false
			}
		case ('A' <= v && v <= 'Z'):
			if !upper {
				v = v + 32
			} else {
				upper = false
			}
		case ('0' <= v && v <= '9'):
			if upper {
				upper = false
			}
		default:
			upper = true
		}
		res += string(v)
	}
	return res
}
