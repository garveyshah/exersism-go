package grains

import (
	"errors"
)

// Calculate the number of grains on a given square
func Square(number int) (uint64, error) {
	var num uint64
	if number < 1 || number > 64 {
		return 0, errors.New("invalid square number")
	}
	num = 1 << (number - 1)
	return num, nil
}

// Calculate the total number of grains on the chessboard
func Total() uint64 {
	var totalGrains uint64
	for square := 1; square <= 64; square++ {
		totalGrains += 1 << (square - 1)
	}
	return totalGrains
}
