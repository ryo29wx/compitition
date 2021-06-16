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
	N := nextInt()
	var A []int
	var B []int
	var C []int
	count := make(map[int]int)

	for i := 0; i < N; i++ {
		A = append(A, nextInt())
		count[A[i]]++
	}
	for i := 0; i < N; i++ {
		B = append(B, nextInt())
	}
	for i := 0; i < N; i++ {
		C = append(C, nextInt())
	}

	ans := 0
	for i := 0; i < N; i++ {
		ans += count[B[C[i]-1]]
	}

	fmt.Println(ans)
}
