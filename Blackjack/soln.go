package main

import "fmt"

func main() {
	two := "ace"
	card1, card2, dealerCard := "ace", "ace", "jack"
	card3, card4, dealerCard1 := "ace", "king", "ace"
	card5, card6, dealerCard2 := "five", "queen", "ace"

	fmt.Println(ParseCard(two))
	fmt.Println(FirstTurn(card1, card2, dealerCard))
	fmt.Println(FirstTurn(card5, card6, dealerCard2))
	fmt.Println(FirstTurn(card3, card4, dealerCard1))
}

func ParseCard(card string) int {
	Value := 0
	Stack := map[string]int{"ace": 11, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9, "ten": 10, "jack": 10, "queen": 10, "king": 10, "other": 0}

	for key, value := range Stack {
		if card == key {
			Value += value
			return value
		}
	}
	return Value
}

func FirstTurn(card1, card2, dealerCard string) string {
	cardV := ParseCard(card1) + ParseCard(card2)
	dealersCards := ParseCard(dealerCard)

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case cardV >= 17 && cardV <= 20:
		return "S"
	case cardV >= 12 && cardV <= 16 && dealersCards >= 7:
		return "H"
	case cardV <= 11:
		return "H"
	case cardV >= 21 && dealersCards < 10:
		return "W"
	default:
		return "S"
	}
}
