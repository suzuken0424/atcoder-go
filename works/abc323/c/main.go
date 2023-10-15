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

	var testNum []int
	for i := 0; i < m; i++ {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		testNum = append(testNum, a)
	}

	// solveProblems := make(map[int][]string)
	// for i := 1; i <= n; i++ {
	// 	scanner.Scan()
	// 	solveProblems[i] = strings.Split(scanner.Text(), "")
	// }

	solveProblems := make(map[int][]string)
	testResultList := make(map[int]int)
	for i := 0; i < n; i++ {
		testResult := 0

		scanner.Scan()
		s := strings.Split(scanner.Text(), "")
		solveProblems[i] = s
		for key, val := range s {
			if val == "o" {
				testResult = testResult + testNum[key]
			}
		}
		testResultList[i] = testResult + i + 1
	}

	fmt.Println(testNum)
	fmt.Println(testResultList)
	fmt.Println(solveProblems)

	max := 0
	for _, val := range testResultList {
		if max < val {
			max = val
		}
	}

	// for _, val := range testResultList {

	// }
}
