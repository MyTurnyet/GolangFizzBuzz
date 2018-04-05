package main

import "testing"

func TestGetString(t *testing.T) {
	returnVal:=CalculateFizzBuzz(1)
	expected := "1"
	if returnVal != expected {
		t.Errorf("GetString return invalid value: %q is not %q", expected, returnVal)
	}
}

func TestThreeShouldReturnFizz(t *testing.T) {
	returnVal:=CalculateFizzBuzz(3)
	expected := "fizz"
	if returnVal != expected {
		t.Errorf("GetString return invalid value: %q is not %q", expected, returnVal)
	}
}