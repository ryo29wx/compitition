package main

import (
	"bufio"
	"fmt"
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

func main() {
	sc.Split(bufio.ScanWords)
	a := nextInt()
	b := nextInt()
	c := nextInt()

	if c == 0 {
		if a <= b {
			fmt.Println("Aoki")
		} else {
			fmt.Println("Takahashi")

		}
	} else {
		if a < b {
			fmt.Println("Aoki")
		} else {
			fmt.Println("Takahashi")
		}
	}

}
