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
	N  int
	M  int
	A  []int
	B  []int
	dp [110][100010]int
)

const (
	Inf = 1 << 40
	//Inf = 0x3f3f3f3f
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
	N = nextInt()
	M = nextInt()
	for i := 0; i < N; i++ {
		A = append(A, nextInt())
	}
	for i := 0; i < M; i++ {
		B = append(B, nextInt())
	}
	sort.Slice(A, func(i, j int) bool { return A[i] < A[j] })
	sort.Slice(B, func(i, j int) bool { return B[i] < B[j] })

	var ans []int

	for _, a := range A {
		left := -1
		right := M
		tmp := Inf

		for right-left > 1 {
			mid := left + (right-left)/2
			//fmt.Println(mid)
			ans1 := B[mid] - a
			if ans1 < 0 {
				left = mid
				ans1 = -ans1
			} else if ans1 > 0 {
				right = mid - 1
			} else {
				ans = append(ans, 0)
				break
			}

			if tmp > ans1 {
				tmp = ans1
			}
		}
		ans = append(ans, tmp)
	}
	sort.Slice(ans, func(i, j int) bool { return ans[i] < ans[j] })

	fmt.Println(ans[0])

}
