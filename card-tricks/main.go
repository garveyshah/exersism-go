package main

import (
	"fmt"
	"os"
	"strings"

	"exercism/card-tricks/cards"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Isage: go run . <slice of numbers>")
	}

	slice := []int{}
	// index := cards.CustomAtoi(os.Args[2])
	// value := cards.CustomAtoi(os.Args[3])

	for _, num := range strings.Split(os.Args[1], " ") {
		slice = append(slice, cards.CustomAtoi(string(num)))
	}

	// fmt.Println(cards.FavoriteCards(slice))
	// fmt.Println(cards.GetItem(slice))
	// fmt.Println(cards.SetItem(slice, index, value))
	fmt.Println(cards.PrependItems(slice, 5, 6, 7))
}
