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

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	k := nextInt()

	var ans int
	var ini int
	for i := 1; i <= n; i++ {
		ini = i * 100
		for j := 1; j <= k; j++ {
			ans = ans + ini + j
		}
	}
	fmt.Println(ans)
}
