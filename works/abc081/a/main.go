package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

const baseRune = 'a'

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	defer flush()

	s := inputInt()
	var count int
	for i := 0; i < 3; i++ {
		if s%10 == 1 {
			count++
		}
		s = s / 10
	}
	outputResult(count)
}

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

/*
バッファにたまったデータを出力先に書き込む際にエラーが発生した場合に、
そのエラーを検知し、プログラムを停止させる役割を果たします。これにより、
データの書き込みが確実に行われるようになります。
*/
func flush() {
	e := wtr.Flush()
	if e != nil {
		panic(e)
	}
}

// 標準入力の1つの数字を取得する
func inputInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

// 標準入力の2つの数字を取得する
func input2Int() (int, int) {
	return inputInt(), inputInt()
}

// 標準入力の文字列を取得する
func inputString() string {
	sc.Scan()
	return sc.Text()
}

// 文字列をintのsliceに変換する
func stringToIntSlice(s string) []int {
	r := make([]int, len(s))
	for i, v := range s {
		r[i] = int(v - baseRune)
	}
	return r
}

// intのsliceを作成する
func createIntSlice(length int, defaultNum int) []int {
	sl := make([]int, length)
	for i := 0; i < length; i++ {
		sl[i] = defaultNum
	}
	return sl
}

// 結果を出力
func outputResult(v ...interface{}) {
	_, e := fmt.Fprintln(wtr, v...)
	if e != nil {
		panic(e)
	}
}
