package main

import (
	"fmt"
	"os"
)

func main() {
	var order bool
	var insert string
	args := os.Args[1:]
	if len(args) == 0 {
		printHelp()
		return
	}
	for len(args) > 0 && len(args[0]) > 0 && args[0][0] == '-' {
		arg := getArg(&args)
		splitIndex := findRune(arg, '=')
		flag := arg[:splitIndex]
		switch flag {
		case "--order", "-o":
			order = true
		case "--insert", "-i":
			if arg > flag { // Some (probably unnecessary) error checking
				insert = arg[splitIndex+1:]
			} else if len(args) > 0 {
				insert = getArg(&args)
			} else {
				insert = ""
			}
		case "--help", "-h":
			printHelp()
			return
		default:
			printHelp()
			return
		}
	}
	var final string
	for len(args) > 0 {
		arg := getArg(&args)
		final += arg
	}
	final += insert
	if order {
		final = sortString(final)
	}
	fmt.Println(final)
}

func getArg(args *[]string) string {
	arg := (*args)[0]
	*args = (*args)[1:]
	return arg
}

func findRune(s string, r rune) int {
	slc := []rune(s)
	var index int = -1
	for i, v := range slc {
		index = i
		if v == r {
			break
		}
		index++ // Increase by 1 if rune not found, so the return == len(s)
	}
	return index
}

func sortString(s string) string {
	sRunes := []rune(s)
	for {
		ready := true
		for i := 0; i < len(sRunes)-1; i++ {
			if sRunes[i] > sRunes[i+1] {
				ready = false
				buffer := sRunes[i]
				sRunes[i] = sRunes[i+1]
				sRunes[i+1] = buffer
			}
		}
		if ready {
			break
		}
	}
	return string(sRunes)
}

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("	 This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("	 This flag will behave like a boolean, if it is called it will order the argument.")
}
