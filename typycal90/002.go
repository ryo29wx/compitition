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

	var a uint
	var b uint

	a = 10
	b = 12

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(n)

	for bits := 0; bits < (1 << uint64(n)); bits++ {
		for i := 0; i < n; i++ {
			if (bits>>uint64(i))&1 == 1 {
				fmt.Println(bits)
			}
		}
	}

}
