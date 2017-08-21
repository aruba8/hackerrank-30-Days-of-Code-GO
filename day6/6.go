package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	var arr []string
	for i := 0; i < n; i++ {
		var str string
		fmt.Scanln(&str)
		arr = append(arr, str)
	}

	for _, v := range arr {
		strs := strings.Split(v, "")
		odd := ""
		even := ""
		for i, k := range strs {
			if i%2 == 0 {
				odd = odd + k
			} else {
				even = even + k
			}
		}
		fmt.Println(odd + " " + even)
	}
}
