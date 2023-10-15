// https://atcoder.jp/contests/abc314/tasks/abc314_c
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Split(bufio.ScanWords)

	var n, m int
	fmt.Scan(&n, &m)

	var s string
	fmt.Scan(&s)

	var c []int
	mp := make(map[int][]string)
	for i := 0; i < n; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		mp[num] = append(mp[num], s[i:i+1])
		c = append(c, num)
	}

	for i := 1; i <= m; i++ {
		lastStr := mp[i][len(mp[i])-1]
		mp[i] = mp[i][:len(mp[i])-1]
		mp[i] = append([]string{lastStr}, mp[i]...)
	}

	res := strings.Builder{}
	for _, val := range c {
		firstStr := mp[val][0]
		mp[val] = mp[val][1:]
		res.WriteString(firstStr)
	}

	fmt.Println(res.String())
}
