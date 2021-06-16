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
	var a int
	var b int
	var c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	if c%2 == 0 {
		aAbs := math.Abs(float64(a))
		bAbs := math.Abs(float64(b))
		if aAbs > bAbs {
			fmt.Println(">")
		} else if aAbs == bAbs {
			fmt.Println("=")
		} else {
			fmt.Println("<")
		}
	} else {
		if a > b {
			fmt.Println(">")
		} else if a == b {
			fmt.Println("=")
		} else {
			fmt.Println("<")
		}
	}
}
