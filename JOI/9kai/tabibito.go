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
	N     int
	M     int
	slice []int
	sum   []int
)

func main() {
	sc.Split(bufio.ScanWords)
	N = nextInt()
	M = nextInt()

	sum = append(sum, 0)
	for i := 1; i < N; i++ {
		tmp := nextInt()
		slice = append(slice, tmp)
		sum = append(sum, sum[i-1]+tmp)
	}

	var move []int
	for i := 1; i <= M; i++ {
		fmt.Println("this1")
		move = append(move, nextInt())
		fmt.Println("this2")
	}

	fmt.Println(move)
	ans := 0
	now := 0

	for _, next := range move {
		if next >= 0 {
			ans = ans + (sum[next+now] - sum[now])
		} else {
			ans = ans + (sum[now] - sum[next+now])
		}
		now = now + next
	}

	fmt.Println(ans)
}
