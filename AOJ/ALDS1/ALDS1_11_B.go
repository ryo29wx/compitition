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
	N int
	K int
	// グラフ探索用
	G          [110][]int
	Visited    [110]bool
	FirstOrder [110]int
	LastOrder  [110]int
	// 動的計画法用
	dp []int
	// 二分探索
	S []int
)

const (
	// Inf = 1 << 20
	Inf = 0x3f3f3f3f
)

func main() {
	sc.Split(bufio.ScanWords)
	///////////
	// 入力
	//////////
	N = nextInt()
	for i := 0; i < N; i++ {
		start := nextInt()
		K = nextInt()
		for j := 0; j < K; j++ {
			G[start] = append(G[start], nextInt()) //隣接nodeの情報を持つ
		}
		// G[b] = append(G[b], a) // 無向グラフ.有向グラフの時はこの行をコメントアウト
	}
	// fmt.Println("G : ", G)

	firstPtr := 0 //lastPtr := 0, 0
	dfs(1, &firstPtr)
	for i := 2; i <= N; i++ {
		if FirstOrder[i] == 0 || LastOrder[i] == 0 {
			dfs(i, &firstPtr)
		}
	}

	for i := 1; i <= N; i++ {
		fmt.Printf("%v ", i)
		fmt.Printf("%v ", FirstOrder[i])
		fmt.Printf("%v ", LastOrder[i])
		fmt.Printf("\n")
	}
}

// 入力用
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

// ダイクストラ用
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

// 動的計画法
// 最大化問題
func chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}

// 最小化問題
func chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}

// 最大公約数 (a,b の最大公約数)
func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

// 二分探索
func binarySearch(key int) int {
	left := -1
	right := len(S)
	for right-left > 1 {
		mid := left + (right-left)/2

		if isOK(mid, key) {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}

func isOK(index, key int) bool {
	if S[index] >= key {
		return true
	} else {
		return false
	}
}

// 深さ優先探索
func dfs(v int, firstPtr *int) {
	*firstPtr++
	FirstOrder[v] = *firstPtr

	Visited[v] = true
	for _, nextV := range G[v] {
		if Visited[nextV] == true {
			continue
		}
		dfs(nextV, firstPtr)
	}
	// fmt.Printf("%v ", v)
	*firstPtr++
	LastOrder[v] = *firstPtr
}
