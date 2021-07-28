package main

import (
	"bufio"
	"container/heap"
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
	g    [500000][]edge //Graph
	r    int
	v    int
	dist [500000]int
)

const (
	INF = 1 << 20
)

func main() {
	sc.Split(bufio.ScanWords)
	///////////
	// 入力
	//////////
	v = nextInt()
	e := nextInt()
	r = nextInt() //start

	for i := 0; i < e; i++ {
		from := nextInt()
		edges := g[from]
		var tmp edge
		tmp.to = nextInt()
		tmp.cost = nextInt()
		edges = append(edges, tmp)
		g[from] = edges
	}

	dijkstra()

}

func dijkstra() {
	for i := 0; i < 500000; i++ {
		dist[i] = INF
	}
	dist[r] = 0
	pq := make(PriorityQueue, 0)
	for _, e := range g[r] {
		edge1 := &edge{
			to:   e.to,
			cost: e.cost,
		}
		heap.Push(&pq, edge1)
	}

	for pq.Len() > 0 {
		e := heap.Pop(&pq).(*edge)
		if dist[e.to] != INF {
			continue
		}
		dist[e.to] = e.cost
		for _, e1 := range g[e.to] {
			edge1 := &edge{
				to:   e1.to,
				cost: e1.cost + e.cost,
			}
			heap.Push(&pq, edge1)
		}
	}

	/////////
	// 出力
	/////////
	for i := 0; i < v; i++ {
		if dist[i] == INF {
			fmt.Println("INF")
		} else {
			fmt.Println(dist[i])
		}
	}

}
