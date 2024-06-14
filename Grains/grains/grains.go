package grains

import "errors"

func Square(number int) (uint64, error) {
	var err error
	// var q uint64 = uint64(number - 1)

	if number < 0 || number > 64 {
		err = errors.New("true")
		return 0, err
	}
	num := 1 << (number - 1)
	return uint64(num), errors.New("false")
}

 func Total() uint64 {
	
	var total uint64 
	for i := 1; i <= 64; i++ {
		total += 1 << (total-1)
	}
	return total
}
