package main

import (
	"flag"
	"fmt"
)

// LeapYear returns if the year is a leap year
func LeapYear(year int) bool {
	if year%4 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	year := flag.Int("year", 0, "Year")
	flag.Parse()

	fmt.Println(LeapYear(*year))
}
