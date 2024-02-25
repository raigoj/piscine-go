package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for { // Bubble sort
		ready := true
		for i := 0; i < len(args)-1; i++ {
			if args[i] > args[i+1] {
				ready = false
				buffer := args[i]
				args[i] = args[i+1]
				args[i+1] = buffer
			}
		}
		if ready {
			break
		}
	}

	for _, s := range args {
		for _, r := range s {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

// rebasepede
