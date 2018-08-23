package main

import "testing"

func TestTrueWhenDivisibleBy4(t *testing.T) {
	if !LeapYear(1996) {
		t.Fail()
	}
}

func TestFalseWhenNotDivisibleBy4(t *testing.T) {
	if LeapYear(1997) {
		t.Fail()
	}
}

func TestTrueWhenDivisibleBy400(t *testing.T) {
	if !LeapYear(1600) {
		t.Fail()
	}
}

func TestFalseWhenDivisibleBy100ButNotBy400(t *testing.T) {
	if LeapYear(1800) {
		t.Fail()
	}
}
