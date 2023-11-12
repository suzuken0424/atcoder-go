package main

import (
	"bufio"
	"os"
)

// スライスの重複削除
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

var scanner = bufio.NewScanner(os.Stdin)

func main() {
}
