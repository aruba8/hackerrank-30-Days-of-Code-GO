package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scanf("%d", &n)

	arr := make([]int, n)

	for l := 0; l < n; l++ {
		fmt.Scanf("%d", &arr[l])
	}

	for k := n - 1; k >= 0; k-- {
		fmt.Printf("%d ", arr[k])
	}

}
