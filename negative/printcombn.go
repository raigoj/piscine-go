package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	var notFirst bool
	s := make([]rune, n)
	CombN(n, s, &notFirst)
	z01.PrintRune('\n')
}

func CombN(depth int, s []rune, notFirst *bool) {
	depth--
	for i := '0'; i <= '9'; i++ {
		s[depth] = i
		if depth == len(s)-1 || s[depth+1] < s[depth] {
			if depth > 0 {
				CombN(depth, s, notFirst)
			} else {
				if *notFirst {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				} else {
					*notFirst = true
				}
				for h := len(s) - 1; h >= 0; h-- {
					z01.PrintRune(s[h])
				}
			}
		}
	}
}
