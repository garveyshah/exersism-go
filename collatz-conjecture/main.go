package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage : go run . <int>")
		return
	}

	num, _ := strconv.Atoi(os.Args[1])

	result, err := CollatzConjecture(num)
	fmt.Println(result, err)
}

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("%d not a positive interger", n)
	}

	var result int

	for n > 1 {
		if n%2 == 0 {
			n = n / 2
			result++
		} else if n%2 != 0 {
			n = ((n * 3) + 1)
			result++
		}
	}
	return result, nil
}
