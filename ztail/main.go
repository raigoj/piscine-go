package main

import (
	"fmt"
	"os"
)

func main() {
	if os.Args[1] != "-c" {
	} else {
		args := os.Args[3:]
		s := Atoi(os.Args[2])
		check := false
		cont := false
		var arr []byte

		for i := 0; i < len(args); i++ {
			cont = false
			file, err := os.Open(args[i])
			fi, e := file.Stat()
			if e != nil {
				cont = true
			}
			if !cont {
				size := fi.Size()
				arr = make([]byte, size)
				// var arr []byte
				file.Read(arr)

			}
			str := string(arr)

			if err != nil {
				fmt.Println(err)
				check = true
				if cont {
					continue
				}

			} else {
				if cont {
					continue
				}
				if len(args) > 1 {
					if i > 0 {
						fmt.Println()
					}
					fmt.Print("==> ")
					fmt.Print(args[i])
					fmt.Print(" <==")
					fmt.Println()
				}
				if len(str) < s {
					fmt.Print(string(arr))
				} else {
					fmt.Print(string(arr[len(string(arr))-s:]))
				}

				file.Close()

			}

		}
		if check {
			os.Exit(1)
		}

	}
}

func Atoi(s string) int {
	arr := []rune(s)
	num := 0
	t := 1
	invchars := false
	negative := false

	for i := len(arr); i > 0; i-- {
		if i == 1 {
			if arr[i-1] == 45 {
				negative = true
			} else if arr[i-1] == 43 {
				invchars = true
			} else {

				num += int(arr[i-1]-'0') * t
				t = t * 10
			}
		} else if arr[i-1] > 57 || arr[i-1] < 48 {
			invchars = true
		} else {

			num += int(arr[i-1]-'0') * t
			t = t * 10
		}
	}
	if invchars {
		num = 0
	}
	if negative {
		num = (num * -1)
	}
	return num
}
