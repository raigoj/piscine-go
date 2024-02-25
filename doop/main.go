package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	a := atoi(args[0])
	b := atoi(args[2])
	op := args[1]
	var res int

	switch op {
	case "+":
		res = a + b
		if (res > a) != (b > 0) {
			return
		}
	case "-":
		res = a - b
		if (res < a) != (b > 0) {
			return
		}
	case "*":
		res = a * b
		if !((res < 0) == ((a < 0) != (b < 0))) || !(res/b == a) {
			return
		}
	case "/":
		if b == 0 {
			printStr("No division by 0")
			return
		}
		res = a / b
	case "%":
		if b == 0 {
			printStr("No modulo by 0")
			return
		}
		res = a % b
	default:
		return
	}

	str := itoa(res)
	if str != "" {
		printStr(str)
	}
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var mul int = 1
	var str []rune
	if n < 0 {
		mul = -1
	}
	for n != 0 {
		r := rune('0' + (n % 10 * mul))
		str = append([]rune{r}, str...)
		n /= 10
	}
	if mul == -1 {
		str = append([]rune{'-'}, str...)
	}
	return string(str)
}

func printStr(str string) {
	os.Stdout.WriteString(str)
	os.Stdout.WriteString("\n")
}

func atoi(s string) int {
	mult := 1
	if s[0] == '-' {
		mult = -1
		s = s[1:]
	}
	var total int
	for _, v := range s {
		if !('0' <= v && v <= '9') {
			total = 0
			os.Exit(0)
		}
		n := int(v - '0')
		total *= 10
		total += n * mult
	}
	if (mult == -1 && total > 0) || (mult == 1 && total < 0) {
		os.Exit(0)
	}
	return total
}
