package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
	V       int
	E       int
	Ans     []int
	G       [110][]edge //vector<vector<int>>
	DP      [110][110]int
	Visited [200100]bool
	Counter [200100]int
)

const (
	Inf = 1 << 20
	//Inf = 0x3f3f3f3f
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
	V = nextInt()
	E = nextInt()
	for i := 0; i < E; i++ {
		s := nextInt() // start
		t := nextInt() // to
		d := nextInt() // distance
		e := edge{t, d}

		G[s] = append(G[s], e) //隣接nodeの情報を持つ
		// G[b] = append(G[b], a)
	}

	for i := 0; i < 110; i++ {
		for j := 0; j < 110; j++ {
			DP[i][j] = Inf
		}
		// 同じ頂点は0
		DP[i][i] = 0
	}

	if warshall() {
		for i := 0; i < V; i++ {
			for j := 0; j < V; j++ {
				if DP[i][j] == Inf {
					fmt.Printf("%s ", "INF")
				} else {
					fmt.Printf("%v ", DP[i][j])
				}
			}
			fmt.Printf("\n")
		}
	}
}

func warshall() bool {
	for start, edges := range G {
		for _, edg := range edges {
			DP[start][edg.to] = edg.cost
		}
	}

	for i := 0; i < V; i++ {
		for j := 0; j < V; j++ {
			for k := 0; k < V; k++ {
				chmin(&DP[j][k], DP[j][i]+DP[i][k])
			}
			if DP[j][j] < 0 {
				fmt.Println("NEGATIVE CYCLE")
				return false
			}
		}
	}
	return true
}
