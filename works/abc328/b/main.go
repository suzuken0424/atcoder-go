package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sliceUniq[T comparable](arr []T) []T {
	unique := make([]T, 0, len(arr))
	seen := make(map[T]struct{})

	for _, v := range arr {
		if _, ok := seen[v]; !ok {
			unique = append(unique, v)
			seen[v] = struct{}{}
		}
	}

	return unique
}

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	var n int
	fmt.Scan(&n)

	var count int
	scanner.Scan()
	dList := strings.Split(scanner.Text(), " ")

	for i, d := range dList {
		index := i + 1
		// 月がゾロ目かどうか
		indexString := strconv.Itoa(index)
		indexStringList := strings.Split(indexString, "")
		indexStringUniqueList := sliceUniq[string](indexStringList)
		if len(indexStringUniqueList) != 1 {
			continue
		}

		min := index % 10
		dNum, _ := strconv.Atoi(d)
		days := index % 10
		for days <= dNum {
			count++
			days = days*10 + min
		}
	}

	fmt.Println(count)
}
