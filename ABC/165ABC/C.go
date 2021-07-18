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
	var dfs func(i, n int)

	dfs = func(i, n int) {
		if n == N-1 {
			A[n] = i
			tmp := 0
			for j := 0; j < Q; j++ {
				if A[b[j]-1]-A[a[j]-1] == c[j] {
					tmp += d[j]
				}
			}
			if ans < tmp {
				ans = tmp
			}
			return
		}

		A[n] = i
		for j := i; j <= M; j++ {
			dfs(j, n+1)
		}
	}

	// 数列の初期値
	for i := 1; i <= M; i++ {
		dfs(i, 0)
	}
	fmt.Println(ans)
}
