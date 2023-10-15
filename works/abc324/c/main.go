package main

// todo テストケースが通らない部分があるので、要確認

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	var n int
	fmt.Scan(&n)

	scanner.Scan()
	t := scanner.Text()

	ans := make([]string, 0)
	for i := 1; i <= n; i++ {
		scanner.Scan()
		s := scanner.Text()
		if s == t {
			ans = append(ans, strconv.Itoa(i))
			continue
		}
		if isChage(s, t) {
			ans = append(ans, strconv.Itoa(i))
			continue
		}
		if isInsert(s, t) {
			ans = append(ans, strconv.Itoa(i))
			continue
		}
		if isDelete(s, t) {
			ans = append(ans, strconv.Itoa(i))
			continue
		}
	}

	fmt.Println(len(ans))
	fmt.Println(strings.Join(ans, " "))
}

func isInsert(s, t string) bool {
	if len(s) != len(t)+1 {
		return false
	}

	var i, j, failCount int
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			i++
			failCount++
		}
	}

	// 最後尾に追加されている可能性があるので、<=が必要
	return failCount <= 1
}

func isChage(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	failCount := 0
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			failCount += 1
		}
	}

	return failCount == 1
}

func isDelete(s, t string) bool {
	if len(s)+1 != len(t) {
		return false
	}

	var i, j, failCount int
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			// 大きい方だけ増やすことで、ズレを補正する
			j++
			failCount++
		}
	}

	// 最後尾が削除されている可能性があるので、<=が必要
	return failCount <= 1
}
