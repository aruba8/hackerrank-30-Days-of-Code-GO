package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var binStack []int

	for n > 0 {
		rm := n % 2
		n = n / 2
		binStack = append(binStack, rm)
	}

	maxCount := 0
	count := 0
	for _, v := range binStack {
		if v == 1 {
			count++
			if count > maxCount {
				maxCount = count
			}
		}
		if v == 0 {
			count = 0
		}
	}

	fmt.Println(maxCount)

}
