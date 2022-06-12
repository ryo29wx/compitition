package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"sort"
)

var sc = bufio.NewScanner(os.Stdin)


var (
	N  int
	H  [100010]int
	dp [200010]int
	seen [200010]bool
	order []int
	G  [200010][]int
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

	// fmt.Println(G)

	for v:=1;v<=N;v++ {
		if seen[v] {
			continue
		}
		rec(v)
	}

	var orderRev []int
	for i:=len(order)-1;i>=0;i-- {
		orderRev = append(orderRev, order[i])
	}
	// fmt.Println(orderRev)

	for _, i := range orderRev {
		// Graph[]が持っている順番にDPを更新していく
		for _, j := range G[i] {
			chmax(&dp[j], dp[i] + 1)
		}
	}

	// fmt.Println(dp)
	ans := 0
	for _ , a := range dp {
		if a > ans {
			ans = a
		}
	}
	fmt.Println(ans)

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

func maxInt(a []int) int {
    sort.Sort(sort.IntSlice(a))
    return a[len(a)-1]
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