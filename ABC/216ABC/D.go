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

var (
	N int
	M int
)

func main() {
	sc.Split(bufio.ScanWords)
	N = nextInt()
	M = nextInt()
	var qSlice [200010][]int
	var cnt [200010][]int
	var q []int

	for i := 0; i < M; i++ {
		k := nextInt()
		var s []int
		for j := 0; j < k; j++ {
			s = append(s, nextInt())
		}
		qSlice[i] = s
		cnt[s[0]] = append(cnt[s[0]], i)
	}

	for i := 0; i <= N; i++ {
		if len(cnt[i]) == 2 {
			q = append(q, i)
		}
	}

	for len(q) != 0 {
		now := q[0]
		q = q[1:]

		for _, p := range cnt[now] {
			qSlice[p] = qSlice[p][1:]
			if len(qSlice[p]) != 0 {
				cnt[qSlice[p][0]] = append(cnt[qSlice[p][0]], p)
				if len(cnt[qSlice[p][0]]) == 2 {
					q = append(q, qSlice[p][0])
				}
			}
		}
	}

	flg := true
	for _, p := range qSlice {
		if len(p) != 0 {
			fmt.Println("No")
			flg = false
			break
		}
	}

	if flg {
		fmt.Println("Yes")
	}

}

func chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}
