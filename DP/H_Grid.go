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
)

const (
	// Inf = 1 << 20
	Inf = 0x3f3f3f3f
	mod = 1_000_000_000 + 7
)



func main() {
	sc.Split(bufio.ScanWords)
	h := nextInt()
	w := nextInt()

	g := make([][]rune, h)
	dp := make([][]int, h)

	for i:=0;i<h;i++ {
		dp[i] = make([]int, w)
	}
	//fmt.Println(dp)

	for i := 0; i < h; i++ {
		line := []rune(next())
		g[i] = line
	}

	for hh, line := range g {
		// fmt.Println(hh)
		// fmt.Println(dp)
		dpLine := make([]int, w)

		for ww, v := range line {
			// fmt.Println(ww, dpLine)
			if v == 35 { 
				dpLine[ww] = 0
			} else if hh == 0 && ww == 0 {
				dpLine[ww] = 1
			} else if hh == 0 {
				dpLine[ww] = dpLine[ww-1]
			} else if ww == 0 {
				dpLine[ww] = dp[hh-1][0]
			} else {
				dpLine[ww] = (dp[hh-1][ww] + dpLine[ww-1]) % mod
			}
		}
		dp[hh] = dpLine
		
	}

	fmt.Println((dp[h-1][w-1]))
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