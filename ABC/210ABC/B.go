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
	slice []int
	c     []int
)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	s := next()

	sRune := []rune(s)
	if len(sRune) == n {
	}

	for c, r := range sRune {
		if string(r) == "1" {
			if c%2 == 0 {
				fmt.Println("Takahashi")
			} else {
				fmt.Println("Aoki")
			}
			break
		}
	}

}
