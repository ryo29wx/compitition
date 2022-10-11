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
	l1 := nextInt()
	r1 := nextInt()
	l2 := nextInt()
	r2 := nextInt()

	ans := 0
	if r1 <= l2 {
		ans = 0
	} else if l2 < r1 && l1 <= l2 {
		ans = l2 - r1 
	} else if l2

}
