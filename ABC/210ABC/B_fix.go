package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	if n == n {
	}

	c := strings.Index(s, "1")
	//fmt.Println(c)
	if (c+1)%2 == 1 {
		fmt.Println("Takahashi")
	} else {
		fmt.Println("Aoki")
	}

}
