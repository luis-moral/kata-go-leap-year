package main

import "testing"

func TestLeapYearWhenDivisibleBy4(t *testing.T) {
	if !LeapYear(1996) {
		t.Fail()
	}
}

func TestLeapYearWhenNotDivisibleBy4(t *testing.T) {
	if LeapYear(1997) {
		t.Fail()
	}
}
