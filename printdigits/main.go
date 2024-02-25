package main

import "github.com/01-edu/z01"

func main() {
	var i rune = 48

	for i <= 57 {
		z01.PrintRune(i)
		i++
	}

	z01.PrintRune('\n')
}
