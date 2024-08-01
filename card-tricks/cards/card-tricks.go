package cards

// FavoriteCards returns a slice with the cards 2, 6, and 9 in the order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at a given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index < 0 || index > len(slice) {
		return -1
	}

	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value need to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < 0 || index >= len(slice) {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, value ...int) []int {
	slice = append(value, slice...)
	return slice
}

// RemoveItem removes an Item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index < 0 || index > len(slice) {
		return slice
	}
	slice = append(slice[:index], slice[index+1:]...)
	return slice
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
