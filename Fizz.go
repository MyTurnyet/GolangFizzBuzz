package main


func Fizz(numberToConvert int) string{
	if numberToConvert %3 != 0 {
		return GetString(numberToConvert)
	}
	return "fizz"
}
