package main

import (
	"fmt"
	"time"

	"exercism/animal-magic/chance"

	"golang.org/x/exp/rand"
)

func main() {
	// Seed the random number generator to ensure different results each run.
	rand.Seed(uint64(time.Now().UnixNano()))

	// Test RollAdie
	dieRoll := chance.RollADie()
	fmt.Printf("Rolled a die: %d\n", dieRoll)

	// Test GenerateWandEnergy
	wandEnergy := chance.GenerateWandEnergy()
	fmt.Printf("Generated wand energy: %.2f\n", wandEnergy)

	// Test ShuffleAnimals
	animals := chance.ShuffleAnimals()
	fmt.Println("Shuffled animals:", animals)
}
