package reverse

func Reverse(input string) string {
	result := []rune(input)
	n := len(result)
	for i := 0; i < n/2; i++ {
		// swap characters
		result[i], result[n-1-i] = result[n-i-1], result[i]
	}
	return string(result)
}
