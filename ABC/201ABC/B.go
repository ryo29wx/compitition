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

	var n1 int
	var n2 int
	var ans string
	var n1name string
	for i := 0; i < n; i++ {
		s := next()
		t := nextInt()

		if t > n1 {
			n2 = n1
			n1 = t
			ans = n1name
			n1name = s
		} else if t > n2 {
			n2 = t
			ans = s
		}
	}

	fmt.Println(ans)
}
