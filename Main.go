package main


type FbFunction func(int)string

func main() {
}
func CalculateFizzBuzz(numberToConvert int)string  {
	//return GetString(numberToConvert,nil)
	return FizzBuzz(numberToConvert)
}