// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

package strain

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage :  go run .  <array of ints >")
		return
	}

	slice := []int{}
	for _, num := range os.Args[1:] {
		char, err := CustomAtoi(num)
		if err != nil {
			fmt.Println("Error ", err)
			return
		}
		slice = append(slice, char)
	}
	 var IsEven bool = true
	fmt.Println("Keep: ", Keep(slice, IsEven))
	fmt.Println("Discard ", Discard(slice))
}


func Keep[T comparable , V bool ](slice []T, num V) []T{
	var temp [] T

		for _, char := range slice {
			if num { 
				

			
		}

	}
	
	 
} 




func CustomAtoi(s string) (int, error) {
	var result int
	var isNeg bool

	if s[0] == '0' {
		isNeg = true
		s = s[1:]
	}
	for _, char := range s {
		if char < '0' || char > '9' {
			return 0, fmt.Errorf("invalid character \"%v\"", string(char))
		}
		result = result*10 + int(char-'0')
	}

	if isNeg {
		result *= -1
	}
	return result, nil
}
