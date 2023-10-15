// https://atcoder.jp/contests/abc323/tasks/abc323_b

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

type Person struct {
	Name string
	Age  int
}

type Result struct {
	PersonNum int
	Score     int
}

func main() {
	scanner.Split(bufio.ScanWords)

	var n int
	fmt.Scan(&n)

	var result []Result
	for i := 1; i <= n; i++ {
		scanner.Scan()
		s := scanner.Text()
		result = append(result, Result{PersonNum: i, Score: strings.Count(s, "o")})
	}

	sort.SliceStable(result, func(i, j int) bool { return result[i].Score > result[j].Score })

	var joinList []string
	for _, val := range result {
		joinList = append(joinList, strconv.Itoa(val.PersonNum))
	}

	fmt.Println(strings.Join(joinList, " "))
}
