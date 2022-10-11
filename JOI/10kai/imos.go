package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func next() string {
	sc.Scan()
	return sc.Text()
}
func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}
func nextFloat() float64 {
	i, _ := strconv.ParseFloat(next(), 64)
	return i
}

func isInteger(x float64) bool {
	return math.Floor(x) == x
}

var (
	H     int
	W     int
	field [10][10]int
	sum   [10][10]int
)

func main() {
	sc.Split(bufio.ScanWords)
	T := nextInt() //閉店時刻
	C := nextInt() //お客さんの数

	var S, E []int
	for i := 0; i < C; i++ {
		S = append(S, nextInt())
		E = append(E, nextInt())
	}

	var table [100]int
	for i := 0; i < T; i++ {
		table[i] = 0
	}

	for i := 0; i < C; i++ {
		table[S[i]]++ // 入店処理: カウントを 1 増やす
		table[E[i]]-- // 出店処理: カウントを 1 減らす
	}

	// シミュレート
	for i := 0; i < T; i++ {
		if 0 < i {
			table[i] += table[i-1]
		}
	}
	// 最大値を調べる
	M := 0
	for i := 0; i < T; i++ {
		if M < table[i] {
			M = table[i]
		}
	}

	fmt.Println(M)

}
