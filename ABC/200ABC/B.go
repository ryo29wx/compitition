/*package main

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
	K := nextInt()

	for i := 0; i < K; i++ {
		if N%200 == 0 {
			N = N / 200
		} else {
			N, _ = strconv.Atoi(strconv.Itoa(N) + "200")
		}
	}

	fmt.Println(N)

	200000
}
