package main

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
	// グラフ探索用
	N          int
	K          int
	G          [110][]int
	Visited    [110]bool
	FirstOrder [110]int
	LastOrder  [110]int
	// 動的計画法用
	// dp     [100009][8]int
	dp     [11][8]int
	Weight [110]int
	Value  [110]int
	//dp     [45]int
	// 二分探索
	S []int
	D []int
)

const (
	// Inf = 1 << 20
	Inf = 0x3f3f3f3f
	mod = 1000000007
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
	var A []int
	var B []int
	var C []int

	for i := 0; i < N; i++ {
		A = append(A, nextInt())
	}
	for i := 0; i < N; i++ {
		B = append(B, nextInt())
	}
	for i := 0; i < N; i++ {
		C = append(C, nextInt())
	}
	sort.Slice(A, func(i, j int) bool { return A[i] < A[j] })
	sort.Slice(C, func(i, j int) bool { return C[i] < C[j] })

	ans := 0
	for i := 0; i < len(B); i++ {
		target := B[i]
		a := binarySearch(target, &A, false)
		c := binarySearch(target, &C, true)
		ans = ans + (a * (len(C) - c))
	}

	fmt.Println(ans)

}

func binarySearch(key int, s *[]int, flg bool) int {
	left := -1       //「index = 0」が条件を満たすこともあるので、初期値は -1
	right := len(*s) // 「index = a.size()-1」が条件を満たさないこともあるので、初期値は a.size()

	/* どんな二分探索でもここの書き方を変えずにできる！ */
	for right-left > 1 {
		mid := left + (right-left)/2

		if flg {
			if isOK(mid, key, s) {
				right = mid
			} else {
				left = mid
			}
		} else {
			if isOK2(mid, key, s) {
				right = mid
			} else {
				left = mid
			}
		}
	}

	/* left は条件を満たさない最大の値、right は条件を満たす最小の値になっている */
	return right
}

func isOK(index, key int, s *[]int) bool {
	if (*s)[index] > key {
		return true
	} else {
		return false
	}
}

func isOK2(index, key int, s *[]int) bool {
	if (*s)[index] >= key {
		return true
	} else {
		return false
	}
}

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
