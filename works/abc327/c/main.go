// https://atcoder.jp/contests/abc327/tasks/abc327_c

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

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

func main() {
	a := [][]string{}
	for i := 0; i < 9; i++ {
		scanner.Scan()
		t := strings.Split(scanner.Text(), " ")
		a = append(a, t)
	}

	checkCol := func() bool {
		for row := 0; row < 9; row++ {
			encountered := map[string]bool{}
			for col := 0; col < 9; col++ {
				if !encountered[a[col][row]] {
					encountered[a[col][row]] = true
				} else {
					return false
				}
			}
		}
		return true
	}

	checkRow := func() bool {
		for col := 0; col < 9; col++ {
			encountered := map[string]bool{}
			for row := 0; row < 9; row++ {
				if !encountered[a[col][row]] {
					encountered[a[col][row]] = true
				} else {
					return false
				}
			}
		}
		return true
	}

	checkBlock := func() bool {
		for col := 0; col < 9; col += 3 {
			for row := 0; row < 9; row += 3 {
				encountered := map[string]bool{}
				for blockCol := 0; blockCol < 3; blockCol++ {
					for blockRow := 0; blockRow < 3; blockRow++ {
						if !encountered[a[col+blockCol][row+blockRow]] {
							encountered[a[col+blockCol][row+blockRow]] = true
						} else {
							return false
						}
					}
				}
			}
		}
		return true
	}

	if !checkCol() || !checkRow() || !checkBlock() {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
