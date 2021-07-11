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

// 入力
var (
	N    int
	Dist [1 << 18]int
	G    [1 << 18][]int
)

const (
	Inf = 1 << 29
)

func getdist(start int) {
	//BFS
	for i := 1; i <= N; i++ {
		Dist[i] = Inf
	}

	var q []int
	q = append(q, start)
	Dist[start] = 0

	for len(q) != 0 {
		pos := q[0]
		q = q[1:]
		for _, to := range G[pos] {
			if Dist[to] == Inf {
				Dist[to] = Dist[pos] + 1
				q = append(q, to)
			}
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	N = nextInt()
	for i := 1; i <= N-1; i++ {
		a := nextInt()
		b := nextInt()
		fmt.Println("=======")
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(G[a])
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}

	getdist(1)
	maxn1 := -1
	maxid1 := -1
	for i := 1; i <= N; i++ {
		if maxn1 < Dist[i] {
			maxn1 = Dist[i]
			maxid1 = i
		}
	}

	getdist(maxid1)
	maxn2 := -1.0
	for i := 1; i <= N; i++ {
		maxn2 = math.Max(float64(maxn2), float64(Dist[i]))
	}

	fmt.Println(maxn2 + 1)

}
