package main

type FizzBuzz func(int, FizzBuzz)string

func Fizz(numberToConvert int, nextCall FizzBuzz) string{
	if numberToConvert %3 != 0 {
		return nextCall(numberToConvert,nil)
	}
	return "fizz"
}
