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
	// グラフ探索用
	N          int
	G          [10010][]edge
	Visited    [110]bool
	FirstOrder [110]int
	LastOrder  [110]int
	Ans        int
	// 動的計画法用
	Weight [110]int
	Value  [110]int
	dp     [45]int
)

const (
	// Inf = 1 << 20
	Inf = 0x3f3f3f3f
)

////////////////
// 動的計画法
////////////////

// 最小化問題
func chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}

// 最大化問題
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
	for i := 1; i < N; i++ {
		start := nextInt()
		end := nextInt()
		cost := nextInt()

		edgeStart := G[start]
		e1 := edge{}
		e1.to = end
		e1.cost = cost
		edgeStart = append(edgeStart, e1)

		// 反転
		edgeEnd := G[end]
		e2 := edge{}
		e2.to = start
		e2.cost = cost
		edgeEnd = append(edgeEnd, e2)

		G[start] = edgeStart
		G[end] = edgeEnd
	}

	for i := 1; i <= N; i++ {
		for j := 2; j <= N; j++ {
			if j > i {
				getMaxDist(i, j)
			}
		}
	}
	fmt.Println(Ans)
}

func getMaxDist(start, end int) {

	// Queueの設定
	var q []int
	q = append(q, start)
	maxDist := 0

	// Queueから一つづつ取り出して、最短経路を更新していく
	// 更新したついでに近傍nodeをさらにQueueに積む
	for len(q) != 0 {
		pos := q[0]
		q = q[1:]
		for _, toEdge := range G[pos] {
			if maxDist < toEdge.cost {
				maxDist = toEdge.cost
				fmt.Printf("%v %v %v \n", start, end, maxDist)
			}
			q = append(q, toEdge.to)
		}
		if pos == end {
			break
		}
	}

	Ans = Ans + maxDist

}
