package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print(strings.Repeat(" ", n-i-1))
		fmt.Println(strings.Repeat("#", i+1))
	}
}
