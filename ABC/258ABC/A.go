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
	k := nextInt()

	a := k / 60
	t := k % 60

	if a >= 1 {
		if t < 10 {
			fmt.Printf("22:0%d", t)
		} else {
			fmt.Printf("22:%d", t)
		}
	} else {
		if t < 10 {
			fmt.Printf("21:0%d", t)
		} else {
			fmt.Printf("21:%d", t)
		}
	}


}
