package main

import (
	"fmt"
	"math"
)

func main() {

	// Declare second integer, double, and String variables.
	var cost float64
	var tips int
	var tax int
	// Read and save an integer, double, and String to your variables.
	fmt.Scan(&cost)
	fmt.Scan(&tips)
	fmt.Scan(&tax)

	p := cost / 100
	t := p * float64(tax)

	pi := cost / 100
	ti := pi * float64(tips)

	preTotaltotalCost := cost + ti + t

	s := fmt.Sprintf("The total meal cost is %d dollars.", int(Round(preTotaltotalCost)))

	fmt.Println(s)
}

//Round input
func Round(input float64) float64 {
	if input < 0 {
		return math.Ceil(input - 0.5)
	}
	return math.Floor(input + 0.5)
}
