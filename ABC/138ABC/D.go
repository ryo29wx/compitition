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
	N       int
	x       int
	Seen    [200010]bool
	Counter [200010]int
	G       [200010][]int //vector<vector<int>>
)

//深さ優先探索
func dfs(v, pa int) {
	for _, nextV := range G[v] {
		if pa != nextV {
			Counter[nextV] += Counter[v]
			dfs(nextV, v)
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	N = nextInt()
	Q := nextInt()
	for i := 1; i <= N-1; i++ {
		a := nextInt()
		b := nextInt()
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}

	for i := 1; i <= Q; i++ {
		c := nextInt()
		x = nextInt()
		Counter[c] = Counter[c] + x
	}

	dfs(1, 0)

	for i := 1; i <= N; i++ {
		fmt.Printf("%d ", Counter[i])
	}
	fmt.Printf("\n")

}
