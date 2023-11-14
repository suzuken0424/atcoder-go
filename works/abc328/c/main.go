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

	n, q := input2Int()
	s := stringToIntSlice(inputString())
	count := createIntSlice(n, 0)

	// 累積和の前準備
	for i := 1; i < n; i++ {
		if s[i-1] == s[i] {
			count[i] = count[i-1] + 1
		} else {
			count[i] = count[i-1]
		}
	}
	// 累積和の計算
	for i := 0; i < q; i++ {
		l, r := input2Int()
		outputResult(count[r-1] - count[l-1])
	}
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
