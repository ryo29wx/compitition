package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)


var (
	N  int
	H  [100010]int
	dp [3010][3010]int
	seen [10]bool
	order []int
	G  [10][]int
	counter int
)

const (
	// Inf = 1 << 20
	Inf = 0x3f3f3f3f
)



func main() {
	sc.Split(bufio.ScanWords)
	N = nextInt()
	M := nextInt()
	for i := 1; i <= M; i++ {
		a := nextInt()
		b := nextInt()
		G[a] = append(G[a], b) //隣接nodeの情報を持つ
		// G[b] = append(G[b], a) // 今回は子ノードしか探索できない有向グラフ
	}

	for v:=0;v<N;v++ {
		if seen[v] {
			continue
		}
		rec(v)
	}
	for i := len(order)-1; i>=0; i-- {
		fmt.Println(order[i])
	}
}

func rec(v int ) {
	seen[v] = true
	for _, next := range G[v] {
		counter++
		if seen[next] {
			continue
		}
		rec(next)
	}
	order = append(order, v)
	
}

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

func chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func isInteger(x float64) bool {
	return math.Floor(x) == x
}

type edge struct {
	to   int
	cost int
}

// PriorityQueue 優先度つきキュー
type PriorityQueue []*edge

// Len pqの長さ
func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

// Swap swap
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push pqにつむ
func (pq *PriorityQueue) Push(x interface{}) {
	e := x.(*edge)
	*pq = append(*pq, e)
}

// Pop 先頭要素を取り出す
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	e := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return e
}

func (pq *PriorityQueue) update(e *edge, to int, cost int) {
	e.to = to
	e.cost = cost
}