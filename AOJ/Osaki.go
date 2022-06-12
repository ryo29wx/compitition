package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
)

func main() {
	sc.Split(bufio.ScanWords)
	var ans []int
	for {
		N = nextInt()
		if N == 0 {
			break
		}

		var start []int
		var end []int
		for i := 0; i < N; i++ {
			var tmpS []int
			var tmpE []int
			s := next()
			e := next() //["00:00:00", "11:11:11"]

			s1 := strings.Split(s, ":")
			for _, v := range s1 {
				kk, _ := strconv.Atoi(v)
				tmpS = append(tmpS, int(kk))
			}
			e1 := strings.Split(e, ":")
			for _, v := range e1 {
				kk, _ := strconv.Atoi(v)
				tmpE = append(tmpE, int(kk))
			}

			start = append(start, tmpS[0]*3600+tmpS[1]*60+tmpS[2])
			end = append(end, tmpE[0]*3600+tmpE[1]*60+tmpE[2])
		}

		T := 3600 * 25
		var table [3600 * 25]int
		for i := 0; i < T; i++ {
			table[i] = 0
		}

		for i := 0; i < N; i++ {
			table[start[i]]++
			table[end[i]]--
		}

		// シミュレート
		for i := 0; i < T; i++ {
			if 0 < i {
				table[i] += table[i-1]
			}
		}
		// 最大値を調べる
		M := 0
		for i := 0; i < T; i++ {
			if M < table[i] {
				M = table[i]
			}
		}

		ans = append(ans, M)

	}

	for _, a := range ans {
		fmt.Println(a)
	}

}
