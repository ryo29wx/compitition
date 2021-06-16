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
	n := nextInt()
	var a []int
	var b []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
	}
	for i := 0; i < n; i++ {
		b = append(b, nextInt())
	}

	count := 0
	for i := 1; i < 11; i++ {
		flg := false
		for j := 0; j < n; j++ {
			if i >= a[j] && i <= b[j] {
				fmt.Println("===========")
				fmt.Println(i)
				fmt.Println(j)
				flg = true
				continue
			} else {
				flg = false
				break
			}

			if flg {
				count++
			}
		}
	}

	fmt.Println(count)

}
