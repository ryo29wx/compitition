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
	N := nextInt()
	M := nextInt()
	Q := nextInt()
	var a, b, c, d [50]int
	for i := 0; i < Q; i++ {
		a[i] = nextInt()
		b[i] = nextInt()
		c[i] = nextInt()
		d[i] = nextInt()
	}

	ans := 0
	A := make([]int, N)
	var dfs func(n, i int)
	dfs = func(n, i int) {
		if i == N-1 {
			A[i] = n

			// calculate sum of points
			point := 0
			for j := 0; j < Q; j++ {
				if A[b[j]-1]-A[a[j]-1] == c[j] {
					point += d[j]
				}
			}
			chmax(&ans, point)
			return
		}

		A[i] = n
		for j := n; j <= M; j++ {
			dfs(j, i+1)
		}
	}

	for i := 1; i <= M; i++ {
		dfs(i, 0)
	}
	fmt.Println(ans)
}

func chmax(a *int, b int) int {
	m := max2(*a, b)
	*a = m
	return m
}
func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
