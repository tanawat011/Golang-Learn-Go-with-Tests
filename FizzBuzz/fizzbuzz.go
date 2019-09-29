package FizzBuzz

import "strconv"

const (
	fizzNumber = 3
	buzzNumber = 5
)

// FizzBuzz function that say ...
func FizzBuzz(number int) string {

	if isFizzBuzz(number) {
		return "FizzBuzz"
	}
	if isFizz(number) {
		return "Fizz"
	}
	if isBuzz(number) {
		return "Buzz"
	}

	return strconv.Itoa(number)
}

func isFizz(number int) bool {
	return number % fizzNumber == 0
}

func isBuzz(number int) bool {
	return number % buzzNumber == 0
}

func isFizzBuzz(number int) bool {
	return isFizz(number) && isBuzz(number)
}
