package main

func FizzBuzz(numberToCovert int) string  {
	if numberToCovert %15 != 0 {
		return Buzz(numberToCovert)
	}
	return "fizzbuzz"
}
