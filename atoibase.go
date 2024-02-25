package piscine

func AtoiBase(s string, base string) int {
	if len(base) < 2 {
		return 0
	}
	for i := 0; i < len(base); i++ {
		if base[i] == '-' || base[i] == '+' {
			return 0
		}
	}
	for i := 0; i < len(base); i++ {
		for j := 1; j < len(base); j++ {
			if i != j {
				if base[i] == base[j] {
					return 0
				}
			}
		}
	}
	mult := 1
	answer := 0
	temp := 1
	for i := len(s); i > 0; i-- {

		for j := 0; j < len(base); j++ {
			if s[i-1] == base[j] {
				temp = j
				break
			}
		}

		answer += temp * mult

		mult *= len(base)
	}
	return answer
}
