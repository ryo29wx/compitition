/** package main

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
	dp [][]int
	N  int
)

func main() {
	sc.Split(bufio.ScanWords)
	N = nextInt()

	var ans []string
	for i := 0; i < 121; i++ {
		if N%2 == 1 {
			N = N - 1
			ans = append(ans, "A")
			continue
		}
		if N == 0 {
			break
		}
		ans = append(ans, "B")
		N = N / 2

	}

	for {
		l := len(ans)

		fmt.Printf(ans[l-1])
		ans = ans[0 : l-1]
		if len(ans) == 0 {
			break
		}
	}
	fmt.Printf("\n")

}
