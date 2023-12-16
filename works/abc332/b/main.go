package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

const baseRune = 'a'
const infinite = 1<<62 - 1

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	defer flush()

	k, g, m := input3Int()
	var gMount, mMount int
	for i := 0; i < k; i++ {
		if gMount == g {
			gMount = 0
			continue
		} else if mMount == 0 {
			mMount = m
			continue
		} else {
			gAmari := g - gMount
			if gAmari > mMount {
				gMount = gMount + mMount
				mMount = 0
			} else {
				gMount = g
				mMount = mMount - gAmari
			}
			continue
		}
	}
	outputResult(gMount, mMount)
}

// 浮動小数点数が整数かどうか
func isInteger(num float64) bool {
	return math.Floor(num) == num
}

// intのsliceの結果を出力する
func outputResultsForInt(intSlice []int) {
	r := make([]string, len(intSlice))
	for i, v := range intSlice {
		r[i] = strconv.Itoa(v)
	}
	outputResult(strings.Join(r, " "))
}

// intのsliceを降順にソート
func sortIntSliceDesc(intSlice []int) []int {
	res := make([]int, len(intSlice))
	copy(res, intSlice)
	sort.Sort(sort.Reverse(sort.IntSlice(res)))
	return res
}

// intのsliceを昇順にソート
func sortIntSliceAsc(intSlice []int) []int {
	res := make([]int, len(intSlice))
	copy(res, intSlice)
	sort.Sort(sort.IntSlice(res))
	return res
}

// 小さい方の値を取得する
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 大きい方の値を取得する
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 平方根を取得
func sqrt(num int) float64 {
	return math.Sqrt(float64(num))
}

// intのsliceから最小値を取り出す
func getMinInt(intSlice []int) int {
	sort.Sort(sort.IntSlice(intSlice))
	return intSlice[0]
}

// intのsliceから最大値を取り出す
func getMaxInt(intSlice []int) int {
	sort.Sort(sort.IntSlice(intSlice))
	return intSlice[len(intSlice)-1]
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

// 標準入力の3つの数字を取得する
func input3Int() (int, int, int) {
	return inputInt(), inputInt(), inputInt()
}

// 任意の数の標準入力を受け取り、intのsliceを返す
func inputIntSlice(lenNum int) []int {
	var res []int
	for i := 0; i < lenNum; i++ {
		res = append(res, inputInt())
	}
	return res
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

// 文字列を昇順にソート
func sortStringAsc(s string) string {
	stringSlice := strings.Split(s, "")
	sort.Strings(stringSlice)
	return strings.Join(stringSlice, "")
}

// 文字列を降順にソート
func sortStringDesc(s string) string {
	stringSlice := strings.Split(s, "")
	sort.Sort(sort.Reverse(sort.StringSlice(stringSlice)))
	return strings.Join(stringSlice, "")
}

// 指定の位置の文字を入れ替える
func replaceCharAtIndex(original string, n int, replacement string) string {
	return string(original[:n]) + replacement + string(original[n+1:])
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
