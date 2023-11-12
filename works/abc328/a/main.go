package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	var n int
	var x int

	fmt.Scan(&n)
	fmt.Scan(&x)

	var addList []int
	for i := 0; i < n; i++ {
		var s int
		fmt.Scan(&s)
		if s <= x {
			addList = append(addList, s)
		}
	}

	sum := 0
	for _, val := range addList {
		sum = sum + val
	}

	fmt.Println(sum)
}
