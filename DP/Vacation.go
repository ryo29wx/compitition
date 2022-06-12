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
	N  int
	H  [100010][3]int
	dp [100010][3]int
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
		for j := 0; j < 3; j++ {
			H[i][j] = nextInt()
			dp[i][j] = 0
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if j == k {
					continue
				}
				chmax(&dp[i+1][k], dp[i][j]+H[i][k])
			}
		}

	}

	var ans int
	for i := 0; i < 3; i++ {
		chmax(&ans, dp[N][i])
	}
	fmt.Println(ans)

}
