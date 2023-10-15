// https://atcoder.jp/contests/abc323/tasks/abc323_a

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	s := scanner.Text()

	sList := strings.Split(s, "")

	result := "Yes"
	for i := 1; i < len(sList); i = i + 2 {
		if sList[i] == "1" {
			result = "No"
		}
	}

	fmt.Println(result)
}
