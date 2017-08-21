package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var i uint64 = 4
	var d = 4.0
	var s = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	var i2 int
	var d2 float64
	var s2 string

	// Read and save an integer, double, and String to your variables.
	fmt.Scan(&i2)
	fmt.Scan(&d2)
	for scanner.Scan() {
		s2 = scanner.Text()
		break
	}

	fmt.Println(i2 + int(i))
	fmt.Printf("%.1f \n", d2+d)
	fmt.Println(s + s2)

}
