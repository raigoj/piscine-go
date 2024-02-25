package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	if len(args) < 1 {
		fmt.Println("File name missing")
		return
	}
	data, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Print(string(data))
}
