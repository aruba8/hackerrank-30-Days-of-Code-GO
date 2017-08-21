package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var s int
	fmt.Scan(&s)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", s, i, (s * i))
	}

}
