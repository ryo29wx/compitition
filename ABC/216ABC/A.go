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
	a := nextFloat()

	b := int(a * 10)
	c := b % 100
	y := c % 10
	x := b / 10

	if y <= 2 && y >= 0 {
		fmt.Printf("%v-\n", x)
	} else if y >= 3 && y <= 6 {
		fmt.Printf("%v\n", x)
	} else {
		fmt.Printf("%v+\n", x)
	}
}
