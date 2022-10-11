/**package main

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
	r := nextInt()
	c := nextInt()

	a11 := nextInt()
	a12 := nextInt()
	a21 := nextInt()
	a22 := nextInt()

	if r == 1 {
		if c == 1 {
			fmt.Println(a11)
		} else {
			fmt.Println(a12)
		}
	} else {
		if c == 1 {
			fmt.Println(a21)
		} else {
			fmt.Println(a22)
		}
	}

	

}
