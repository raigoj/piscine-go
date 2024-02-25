package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	f := AtoiBase(nbr, baseFrom)
	return NbrBase(f, baseTo)
}

func NbrBase(nbr int, base string) string {
	if len(base) < 2 {
		return "NV"
	}
	for i := 0; i < len(base); i++ {
		if base[i] == '-' || base[i] == '+' {
			return "NV"
		}
	}
	for i := 0; i < len(base); i++ {
		for j := 1; j < len(base); j++ {
			if i != j {
				if base[i] == base[j] {
					return "NV"
				}
			}
		}
	}
	neg := false
	if nbr < 0 {
		neg = true
	}
	s := ""

	for nbr != 0 {
		if nbr%len(base) == 0 {
			s = string(base[0]) + s
			nbr /= len(base)
			if neg {
				nbr *= -1
			}

		} else {
			temp := nbr % len(base)
			nbr -= temp
			nbr /= len(base)
			if neg {
				nbr *= -1
			}
			if temp < 0 {
				temp *= -1
			}
			s = string(base[temp]) + s

		}
	}
	return s
}
