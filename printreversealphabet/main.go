package main

import "github.com/01-edu/z01"

func main() {
	var i rune = 122

	for i >= 97 {
		z01.PrintRune(i)
		i--
	}

	z01.PrintRune('\n')
}
