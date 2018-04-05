package main

import "testing"

func Test_ReturnStringOfInt(t *testing.T) {
	returnVal:=CalculateFizzBuzz(1)
	expected := "1"
	if returnVal != expected {
		t.Errorf("GetString return invalid value: %q is not %q", expected, returnVal)
	}
}

func Test_MultipleOfThreeShouldReturnFizz(t *testing.T) {
	returnVal:=CalculateFizzBuzz(3)
	expected := "fizz"
	if returnVal != expected {
		t.Errorf("GetString return invalid value: %q is not %q", expected, returnVal)
	}
}
func Test_MultipleOfFiveShouldReturnBuzz(t *testing.T) {
	returnVal:=CalculateFizzBuzz(5)
	expected := "buzz"
	if returnVal != expected {
		t.Errorf("GetString return invalid value: %q is not %q", expected, returnVal)
	}
}

func Test_MultipleOfFifteenShouldReturnBuzz(t *testing.T) {
	returnVal := CalculateFizzBuzz(15)
	expected := "fizzbuzz"
	if returnVal != expected {
		t.Errorf("GetString return invalid value: %q is not %q", expected, returnVal)
	}
}
