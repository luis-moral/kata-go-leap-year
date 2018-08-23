package main

import (
	"flag"
	"fmt"
)

// LeapYear returns if the year is a leap year
func LeapYear(year int) bool {
	return year%4 == 0 && !(year&100 == 0 && year%400 != 0)
}

func main() {
	year := flag.Int("year", 0, "Year")
	flag.Parse()

	fmt.Println(LeapYear(*year))
}
