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
	x := nextInt()
	a := nextInt()
	d := nextInt()
	n := nextInt()

	ans := false
	if x <= a {
		fmt.Println(a-x)
		ans = true
	}

	if !ans && (x >= (a+(d*n))) {
		fmt.Println(x - (a+(d*n)))
		ans = true
	} 

	if !ans {
		amari := (x-a)%d
		// sho := x/d
		if amari <= (d/2) {
			fmt.Println(amari)
		} else {
			fmt.Println(d - amari)
		}
	}
}
