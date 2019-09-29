package FizzBuzz

import "testing"

func TestFizzBuzzWithRightCode(t *testing.T) {

	fizzBuzzTable := map[int]string{
		1: "1",
		2: "2",
		3: "Fizz",
		4: "4",
		5: "Buzz",
		6: "Fizz",
		7: "7",
		8: "8",
		9: "Fizz",
		10: "Buzz",
		11: "11",
		12: "Fizz",
		13: "13",
		14: "14",
		15: "FizzBuzz",
	}

	for number, want := range fizzBuzzTable {
		say := FizzBuzz(number)
		assertFizzBuzz(t, want, say)
	}
}

func assertFizzBuzz(t *testing.T, want, say string) {
	if say != want {
		t.Errorf("Want %s, but said %s", want, say)
	}
}
