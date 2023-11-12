package main

import (
	"fmt"
	"strconv"
)

// 数字がゾロ目かを判定する
func checkMatchingDice(num int) bool {
	onePlace := num % 10
	for ; num > 0; num /= 10 {
		if num%10 != onePlace {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Scan(&n)

	var count int
	// 組み合わせた100*100なので、全通り見る
	for month := 1; month <= n; month++ {
		var d int
		fmt.Scan(&d)

		for day := 1; day <= d; day++ {
			checkNum, _ := strconv.Atoi(strconv.Itoa(month) + strconv.Itoa(day))
			if checkMatchingDice(checkNum) {
				count++
			}

		}
	}

	fmt.Println(count)
}
