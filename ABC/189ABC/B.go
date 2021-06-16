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
	N := nextInt()
	X := nextFloat()

	var intake float64
	for i := 0; i < N; i++ {
		V := nextFloat()
		P := nextFloat()
		intake = intake + (V * P / 100)

		if intake > X {
			fmt.Println(i + 1)
			break
		}
	}

	if intake <= X {
		fmt.Println(-1)
	}

}
