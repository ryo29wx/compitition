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
	H int
	W int
	N int
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

type point struct {
	A int
	B int
}

func main() {
	sc.Split(bufio.ScanWords)
	///////////
	// 入力
	//////////
	fmt.Scan(&H)
	fmt.Scan(&W)
	fmt.Scan(&N)

	var aSlice []int
	var bSlice []int
	aSet := make(map[int]struct{})
	bSet := make(map[int]struct{})

	var pSlice []point
	for i := 1; i <= N; i++ {
		a := nextInt()
		b := nextInt()

		aSet[a] = struct{}{}
		bSet[b] = struct{}{}
		p := point{a, b}
		pSlice = append(pSlice, p)
	}

	for k, _ := range aSet {
		aSlice = append(aSlice, k)
	}
	for k, _ := range bSet {
		bSlice = append(bSlice, k)
	}
	sort.Slice(aSlice, func(i, j int) bool { return aSlice[i] < aSlice[j] })
	sort.Slice(bSlice, func(i, j int) bool { return bSlice[i] < bSlice[j] })

	for _, p := range pSlice {
		a := p.A
		b := p.B
		for k, v := range aSlice {
			if a == v {
				fmt.Printf("%v ", 1+k)
			}
		}
		for k, v := range bSlice {
			if b == v {
				fmt.Printf("%v\n", 1+k)
			}
		}
	}

}
