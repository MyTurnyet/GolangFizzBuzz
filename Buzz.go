package main

func Buzz(numberToConvert int, nextFunction FizzBuzz ) string{
	if numberToConvert != 5 && numberToConvert != 10 {
		return nextFunction(numberToConvert, GetString)
	}
	return "buzz"
}