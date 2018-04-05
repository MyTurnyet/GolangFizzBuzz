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
func Test_FiveShouldReturnBuzz(t *testing.T) {
	returnVal:=CalculateFizzBuzz(5)
	expected := "buzz"
	if returnVal != expected {
		t.Errorf("GetString return invalid value: %q is not %q", expected, returnVal)
	}
}
func Test_TenShouldReturnBuzz(t *testing.T) {
	returnVal:=CalculateFizzBuzz(10)
	expected := "buzz"
	if returnVal != expected {
		t.Errorf("GetString return invalid value: %q is not %q", expected, returnVal)
	}
}