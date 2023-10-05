package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// scannerを作成
var scanner = bufio.NewScanner(os.Stdin)

func next() string {
	scanner.Scan()
	return scanner.Text()
}

func nextInt() int {
	intVal, _ := strconv.Atoi(next())
	return intVal
}

func main() {
	scanner.Split(bufio.ScanWords)
	res := strings.Builder{}

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// [出目][かけ数]:[何番目の人か]
	mp := make(map[int]map[int][]int)
	for i := 0; i < n; i++ {
		scanner.Scan()
		c, _ := strconv.Atoi(scanner.Text())
		for j := 0; j < c; j++ {
			scanner.Scan()
			a, _ := strconv.Atoi(scanner.Text())
			if mp[a] == nil {
				mp[a] = map[int][]int{}
			}
			mp[a][c] = append(mp[a][c], i+1)
		}
	}

	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())

	minKey := 100
	for key, _ := range mp[x] {
		if minKey > key {
			minKey = key
		}
	}

	res.WriteString(strconv.Itoa(len(mp[x][minKey])) + "\n")
	for i := 0; i < len(mp[x][minKey]); i++ {
		val := mp[x][minKey][i]
		if i+1 == len(mp[x][minKey]) {
			res.WriteString(strconv.Itoa(val))
		} else {
			res.WriteString(strconv.Itoa(val) + " ")
		}
	}

	fmt.Println(res.String())
}
