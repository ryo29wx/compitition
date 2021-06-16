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

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	m := nextInt()

	var s []int
	var c []int
	for i := 0; i < m; i++ {
		s = append(s, nextInt())
		c = append(c, nextInt())
	}

	flg := true
	var ans string
	for i := 0; i < 1000; i++ {
		str := strconv.Itoa(i)
		if len(str) != n {
			continue
		}

		ans = str
		flg = false
		for j := 0; j < m; j++ {
			tmp_s := s[j]
			tmp_c := c[j]
			// けた
			if tmp_s > n {
				continue
			}
			dig, _ := strconv.Atoi(str[tmp_s-1 : tmp_s])
			if tmp_c != dig {
				flg = false
				break
			} else {
				flg = true
				continue
			}
		}

		if !flg {
			continue
		} else {
			break
		}
	}

	if !flg {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}

}
