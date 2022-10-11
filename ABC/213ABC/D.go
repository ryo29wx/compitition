/** package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

// 入力用
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

/////////////////////
// 以下ダイクストラ  //
/////////////////////
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

var (
	N       int
	Backlog map[int]int
	G       [200010][]int //vector<vector<int>>
	Visited [200010]bool
	Counter [200010]int
)

const (
	// Inf = 1 << 20
	Inf = 0x3f3f3f3f
)

func chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}

func chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	///////////
	// 入力
	//////////
	N = nextInt()
	for i := 1; i <= N-1; i++ {
		a := nextInt()
		b := nextInt()
		G[a] = append(G[a], b) //隣接nodeの情報を持つ
		G[b] = append(G[b], a)
	}
	for k := 1; k <= N-1; k++ {
		sort.Slice(G[k], func(i, j int) bool { return G[k][i] < G[k][j] })
	}
	Backlog = make(map[int]int, 200010)
	fmt.Printf("%v ", 1)
	dfs(1)

}

func dfs(v int) {
	Visited[v] = true
	for _, nextV := range G[v] {
		if Visited[nextV] == true {
			continue
		}
		Backlog[nextV] = v
		fmt.Printf("%v ", nextV)
		dfs(nextV)
	}
	if v != 1 {
		if val, ok := Backlog[v]; ok {
			fmt.Printf("%v ", val)
		}
	}
}
