/** package main

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
	N  int
	H  [100010]int
	dp [100010]int
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

func main() {
	sc.Split(bufio.ScanWords)
	///////////
	// 入力
	//////////
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		H[i] = nextInt()
	}

	for i := 0; i < 100010; i++ {
		dp[i] = 1 << 40
	}

	dp[0] = 0
	for i := 1; i < N; i++ {
		chmin(&dp[i], dp[i-1]+int(math.Abs(float64(H[i]-H[i-1]))))
		if i > 1 {
			chmin(&dp[i], dp[i-2]+int(math.Abs(float64(H[i]-H[i-2]))))
		}
		//if i == 1 {
		//	dp[i] = min(dp[i], dp[i-1]+int(math.Abs(float64(H[i]-H[i-1]))))
		//} else {
		//	dp[i] = min(dp[i-1]+int(math.Abs(float64(H[i]-H[i-1]))), dp[i-2]+int(math.Abs(float64(H[i]-H[i-2]))))
		//}

	}
	fmt.Println(dp[N-1])

}
