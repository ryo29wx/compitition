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
	p := nextInt()

	tmp1 := 1
	tmp2 := 2 * tmp1
	tmp3 := 3 * tmp2
	tmp4 := 4 * tmp3
	tmp5 := 5 * tmp4
	tmp6 := 6 * tmp5
	tmp7 := 7 * tmp6
	tmp8 := 8 * tmp7
	tmp9 := 9 * tmp8
	tmp10 := 10 * tmp9

	tmp := p / tmp10
	remainder := p % tmp10

	tmp = tmp + (remainder / tmp9)
	remainder = remainder % tmp9

	tmp = tmp + (remainder / tmp8)
	remainder = remainder % tmp8

	tmp = tmp + (remainder / tmp7)
	remainder = remainder % tmp7

	tmp = tmp + (remainder / tmp6)
	remainder = remainder % tmp6

	tmp = tmp + (remainder / tmp5)
	remainder = remainder % tmp5

	tmp = tmp + (remainder / tmp4)
	remainder = remainder % tmp4

	tmp = tmp + (remainder / tmp3)
	remainder = remainder % tmp3

	tmp = tmp + (remainder / tmp2)
	remainder = remainder % tmp2

	tmp = tmp + (remainder / tmp1)
	remainder = remainder % tmp1

	fmt.Println(tmp)
}
