package main

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
	dp [3010][3010]int
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	sc.Split(bufio.ScanWords)
	///////////
	// 入力
	//////////
	s := next()
	t := next()
	
	sRune := []rune(s)
	tRune := []rune(t)
	// dp[i][j]=(sをi文字目、tをj文字目までみた時の最長共通部分列の長さ)

	for i:=1;i<len(sRune)+1;i++ {
		for j:=1;j<len(tRune)+1;j++ {
			if sRune[i-1] == tRune[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	l := dp[len(sRune)][len(tRune)]
	i := len(sRune) -1
	j := len(tRune) -1

	// fmt.Println(l, i, j)

	var ans []rune
	for l>0 {
		// fmt.Println("dp: ", dp[i][j] )

		if sRune[i] == tRune[j] {
			ans = append(ans, sRune[i])

			// fmt.Println("if:", i, j, l)
			i--
			j--
			l--
		} else if i == 0 {
			// fmt.Println("i==0:", i, j)
			j--
			continue
		} else if dp[i+1][j+1] == dp[i][j+1] {
			// fmt.Println("elseif:", i, j)
			i--
		} else {
			// fmt.Println("else:", i, j)
			j--
		}
	}

	for i=len(ans)-1;i>=0;i-- {
		fmt.Printf(string(ans[i]))
	}
	fmt.Println()

}
