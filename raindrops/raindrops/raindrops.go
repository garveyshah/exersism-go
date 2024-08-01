package raindrops

func Convert(number int) string {
	var result string

	switch {
	case number%3 == 0 && number%5 == 0 && number%7 == 0:
		result += "PlingPlangPlong"
	case number%3 == 0 && number%5 == 0:
		result += "PlingPlang"
	case number%3 == 0 && number%7 == 0:
		result += "PlingPlong"
	case number%5 == 0 && number%7 == 0:
		result += "PlangPlong"
	case number%3 == 0:
		result += "Pling"
	case number%5 == 0:
		result += "Plang"
	case number%7 == 0:
		result += "Plong"
	default:
		result = CustomItoa(number)
	}
	return result
}

// Function to convert interger to string
func CustomItoa(num int) string {
	var result []byte
	var isNeg bool

	if num < 0 {
		isNeg = true
		num *= -1
	}

	if num == 0 {
		return "0"
	}

	for num > 0 {
		r := num % 10
		result = append(result, byte(r+'0'))
		num = num / 10
	}
	if isNeg {
		result = append(result, '-')
	}

	// Reverse the result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}

// Function to convert string to interger.
func CustomAtoi(num string) int {
	var result int
	var isNeg bool

	if num[0] == '-' {
		isNeg = true
		num = num[1:]
	}

	for _, char := range num {
		if char < '0' || char > '9' {
			return 0 //, fmt.Errorf("invalid character %v", string(char))
		}
		result = result*10 + int(char-'0')
	}
	if isNeg {
		result *= -1
	}
	return result //, nil
}
