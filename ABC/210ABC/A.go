/* package main

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
	slice []int
	c     []int
)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := nextInt()
	x := nextInt()
	y := nextInt()

	var ans int
	if n > a {
		ans = a*x + (n-a)*y
	} else {
		ans = n * x
	}

	fmt.Println(ans)
}
