/** package main

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
	N      int
	W      int
	Weight [110]int
	Value  [110]int
	dp     [110][100010]int
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

	n := nextInt()
	var amari int
	x1 := n / 1000
	amari = n % 1000

	x2 := amari / 100
	amari = amari % 100

	x3 := amari / 10
	amari = amari % 10

	x4 := amari

	flg := false
	if x1 == x2 && x2 == x3 && x3 == x4 && x4 == x1 {
		fmt.Println("Weak")
		flg = true
	} else if x1+1 == x2 || (x1 == 9 && x2 == 0) {
		if x2+1 == x3 || (x2 == 9 && x3 == 0) {
			if x3+1 == x4 || (x3 == 9 && x4 == 0) {
				fmt.Println("Weak")
				flg = true
			}
		}
	}

	if !flg {
		fmt.Println("Strong")
	}

}
