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

func main() {
	sc.Split(bufio.ScanWords)
	r := nextFloat()
	x := nextFloat()
	y := nextFloat()

	u := math.Sqrt(math.Pow(x, 2.0) + math.Pow(y, 2.0))

	if isInteger(u) {
		if u/r == 0 {
			fmt.Println(int(u) % r)
			os.Exit(1)
		}
	}

	u / r

}
