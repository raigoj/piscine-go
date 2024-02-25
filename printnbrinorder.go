package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var s []int
	slc, neg := sliceNbr(n, s)
	sortIntegerTable(slc)
	if neg {
		z01.PrintRune('-')
	}
	for _, v := range slc {
		z01.PrintRune(rune(v + '0'))
	}
}

func sliceNbr(n int, slc []int) ([]int, bool) {
	neg := false
	if n < 0 {
		neg = true
		n *= -1
	}

	f := n / 10
	if f != 0 {
		slc, _ = sliceNbr(f, slc)
	}
	k := n % 10
	slc = append(slc, k)
	return slc, neg
}

func sortIntegerTable(table []int) {
	for {
		ready := true
		for i := 0; i < len(table)-1; i++ {
			if table[i] > table[i+1] {
				ready = false
				buffer := table[i]
				table[i] = table[i+1]
				table[i+1] = buffer
			}
		}
		if ready {
			break
		}
	}
}

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
