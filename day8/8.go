package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	phones := make(map[string]int)

	for i := 0; i < n; i++ {
		var name string
		var number int
		fmt.Scanf("%s", &name)
		fmt.Scanf("%d", &number)
		phones[name] = number
	}
	for {
		var query string
		_, err := fmt.Scanf("%s", &query)
		if err != nil {
			break
		}
		num, ok := phones[query]

		if ok {
			fmt.Printf("%s=%d\n", query, num)
		} else {
			fmt.Println("Not found")
		}
	}

}
