package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var (
	N int
)

func main() {
	sc.Split(bufio.ScanWords)
	N = nextInt()

	var ansTable [1000010]int
	for i := 0; i < N; i++ {
		ansTable[nextInt()]++
		ansTable[nextInt()+1]--
	}
	// fmt.Println(ansTable)

	// シミュレート
	for i := 0; i < len(ansTable); i++ {
		if 0 < i {
			ansTable[i] += ansTable[i-1]
		}
	}
	// fmt.Println(ansTable)

	// 最大値を調べる
	M := 0
	for i := 0; i < len(ansTable); i++ {
		if M < ansTable[i] {
			M = ansTable[i]
		}
	}

	fmt.Println(M)

}

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
