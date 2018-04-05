package main

import "strconv"

func GetString(numberToCovert int, nextCall FizzBuzz) string {
	return strconv.Itoa(numberToCovert)
}
