package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var upper bool
	args := os.Args[1:]
	if len(args) > 0 {
		if args[0] == "--upper" {
			upper = true
			args = args[1:]
		}
		for _, s := range args {
			i := atoi(s) - 1
			if 0 <= i && i <= 'z'-'a' {
				if upper {
					z01.PrintRune(rune('A' + i))
				} else {
					z01.PrintRune(rune('a' + i))
				}
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func atoi(s string) int {
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
