package main

import (
	"fmt"
	"os"

	"exercism/raindrops/raindrops"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <number>")
	}

	fmt.Println(raindrops.Convert(raindrops.CustomAtoi(os.Args[1])))
}
