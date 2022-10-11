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
	dp []int
	N  int
	K  int
)

func main() {
	sc.Split(bufio.ScanWords)
	N = nextInt()

	for i := 0; i < N; i++ {

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

func chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}
