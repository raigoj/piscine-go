package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, baseString string) {
	base := []rune(baseString)
	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	for i := range base {
		if base[i] == '-' || base[i] == '+' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}

	slc, neg := sliceNbrBase(nbr, len(base))
	if neg {
		z01.PrintRune('-')
	}
	for _, v := range slc {
		if neg {
			v *= -1
		}
		z01.PrintRune(base[v])
	}
}

func sliceNbrBase(num, baseLen int) ([]int, bool) {
	neg := false
	if num < 0 {
		neg = true
	}

	var slc []int
	for {
		slc = prependInt(slc, num%baseLen)
		newNum := num / baseLen
		if newNum == 0 {
			break
		}
		num = newNum
	}
	return slc, neg
}

func prependInt(x []int, y int) []int {
	x = append(x, 0)
	copy(x[1:], x)
	x[0] = y
	return x
}
