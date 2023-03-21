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

var (
	N, L, K int
	A       []int
)

func main() {
	sc.Split(bufio.ScanWords)
	N = nextInt()
	L = nextInt()
	K = nextInt()

	for i := 0; i < N; i++ {
		A = append(A, nextInt())
	}

	// 1. めぐる式二分探索
	high := L + 1
	low := -1
	for high-low > 1 {
		mid := low + (high-low)/2
		if solve(mid) == false {
			high = mid
		} else {
			low = mid
		}
	}

	fmt.Println(low)
}

func solve(mid int) bool {

	count := 0
	tmp := 0
	for i := 0; i < N; i++ {
		if A[i]-tmp >= mid && L-A[i] >= mid {
			count++
			tmp = A[i]
		}
	}

	if count >= K {
		return true
	}
	return false

}
