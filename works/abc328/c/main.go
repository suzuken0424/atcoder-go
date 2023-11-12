package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	var n int
	var q int
	var s string

	fmt.Scanf("%d %d", &n, &q)
	fmt.Scan(&s)

	for i := 0; i < q; i++ {
		var l int
		var r int
		fmt.Scanf("%d %d", &l, &r)

		var count int
		for j := l - 1; j < r-1; j++ {
			if s[j] == s[j+1] {
				count++
			}
		}
		fmt.Println(count)
	}
}
