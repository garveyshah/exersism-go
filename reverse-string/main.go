package main

import (
	"fmt"
	"os"

	"exercism/reverse-string/reverse"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <string>")
	}

	ascii1 := reverse.Reverse(os.Args[1])
	fmt.Println(ascii1)
}
