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
	x := nextInt()

	ans := 0
	for i := 1; i <= n; i++ {
		a := nextInt()
		if i%2 == 0 {
			ans = ans + a - 1
		} else {
			ans = ans + a
		}
	}

	if ans > x {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
