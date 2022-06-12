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
	H int
	W int
	// グラフ探索用
	G          [55][55]int
	Visited    [55][55]bool
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

	//  一旦 0 で初期化しておく
	for i := 0; i < 55; i++ {
		for j := 0; j < 55; j++ {
			G[i][j] = 0
		}
	}

	var Ans []int
	for {
		W = nextInt()
		H = nextInt()
		if W == 0 || H == 0 {
			break
		}
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				G[i][j] = nextInt()
			}
		}
		ans := 0
		for h := 0; h < H; h++ {
			for w := 0; w < W; w++ {
				if G[h][w] == 0 {
					continue
				}
				dfs(h, w)
				ans++
			}
		}
		Ans = append(Ans, ans)
	}

	for _, v := range Ans {
		fmt.Println(v)
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

type point struct {
	x      int
	y      int
	status int
}

// 深さ優先探索
func dfs(h, w int) {
	G[h][w] = 0

	// 八方向を探索
	for dh := -1; dh <= 1; dh++ {
		for dw := -1; dw <= 1; dw++ {
			nh := h + dh
			nw := w + dw

			// 場外アウトしたり、0 だったりはスルー
			if nh < 0 || nh >= H || nw < 0 || nw >= W {
				continue
			}
			if G[nh][nw] == 0 {
				continue
			}
			// 再帰的に探索
			dfs(nh, nw)
		}
	}
}
