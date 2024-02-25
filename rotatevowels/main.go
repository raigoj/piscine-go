package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	var str string
	for len(args) > 0 {
		str += getArg(&args)
		str += " "
	}

	slc := []rune(str)
	newSlc := dupSlice(slc)

	vowIdx := getVowelIndexes(str)
	vowLen := len(vowIdx) - 1
	for i, v := range vowIdx {
		iNew := vowLen - i
		newSlc[vowIdx[iNew]] = slc[v]
	}

	for _, r := range newSlc {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func getArg(args *[]string) string {
	arg := (*args)[0]
	*args = (*args)[1:]
	return arg
}

func getVowelIndexes(str string) []int {
	var iSlice []int
	slc := []rune(str)
	for i, v := range slc {
		if isVowel(v) {
			iSlice = append(iSlice, i)
		}
	}
	return iSlice
}

func isVowel(r rune) bool {
	return false ||
		r == 'a' || r == 'A' ||
		r == 'i' || r == 'I' ||
		r == 'u' || r == 'U' ||
		r == 'e' || r == 'E' ||
		r == 'o' || r == 'O'
}

func dupSlice(slc []rune) []rune {
	newSlc := make([]rune, len(slc))
	copy(newSlc, slc)
	return newSlc
}
