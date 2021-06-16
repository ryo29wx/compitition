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

	max4 := n / 4
	max7 := n / 7
	flg := false
	for i := 0; i <= max4; i++ {
		for j := 0; j <= max7; j++ {
			if 4*i+7*j == n {
				flg = true
				break
			}
		}
	}

	if flg {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")

	}
}
