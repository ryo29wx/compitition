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

func main() {
	sc.Split(bufio.ScanWords)
	///////////
	// 入力
	//////////

	n := nextInt()
	m := nextInt()
	G := make([][]int, n+1)
	rG := make([][]int, n+1)

	// var G [][]int
	// var rG [][]int
	for i:=0;i<m;i++ {
		a := nextInt()
		b := nextInt()

		G[a] = append(G[a], b)
		rG[b] = append(rG[b], a) //辺を反転させたグラフをあらかじめ作っておく
	 }
	 
	 timestamp := make([]int, n+1)
	 visited := make([]bool, n+1)
	 lastT := 0

	 for i:=1;i<=n;i++ {
		if !visited[i] {
			dfs(i, &lastT, &G, &visited, &timestamp)
		}
	 }

	 fmt.Println(timestamp)
	 fmt.Println(visited)


	 var groups [][]int
	 for i:=1;i<=n;i++ {
		max := findMax(&timestamp, &visited)
		if max == 0 {
			break
		}
		group := make([]int, 0)
		dfs2(max, &rG, &visited, &group)

		groups = append(groups, group)
	 }

	 fmt.Println(groups)
}

func dfs(i int, lastT *int, G *[][]int , visited *[]bool, timestamp *[]int, ) {
	if (*visited)[i] {
		return 
	}
	(*visited)[i] = true

	for _, nextV := range (*G)[i] {
		dfs(nextV, lastT, G, visited, timestamp)
	}
	(*lastT)++
	
	(*timestamp)[i] = *lastT

}

func dfs2(i int, G *[][]int , visited *[]bool, group *[]int ) {
	if !(*visited)[i] {
		return 
	}
	fmt.Println(i, visited, G)

	(*visited)[i] = false
	*group = append(*group, i)

	for _, nextV := range (*G)[i] {
		dfs2(nextV, G, visited, group)
	}
}

func findMax(ts *[]int, visited *[]bool) int {
	tmpV := 0
	tmpK := 0
	for k, v := range *ts {
		if !(*visited)[k] {
			continue
		}
		if v > tmpV {
			tmpV = v
			tmpK = k
		}
	}

	return tmpK
}