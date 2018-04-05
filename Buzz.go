package main

func Buzz(numberToConvert int) string{
	if numberToConvert % 5 != 0 {
		return Fizz(numberToConvert)
	}
	return "buzz"
}