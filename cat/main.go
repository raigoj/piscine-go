package main

import (
	"io"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		data, e := io.ReadAll(os.Stdin)
		check(e)
		os.Stdout.Write(data)
		return
	}
	for _, filename := range args {
		data, e := os.ReadFile(filename)
		check(e)
		os.Stdout.Write(data)
	}
}

func check(e error) {
	if e != nil {
		os.Stdout.WriteString("ERROR: ")
		os.Stdout.WriteString(e.Error())
		os.Stdout.WriteString("\n")
		os.Exit(1)
	}
}
