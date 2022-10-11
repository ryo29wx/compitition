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
	k := nextInt()

	var kk []int
	for i:=0; i<k; i++ {
		kk = append(kk, nextInt())
	}

	var xs []int
	var ys []int

	for i:=0; i<n; i++ {
		xs = append(xs, nextInt())
		ys = append(ys, nextInt())
	}


	var ans float64 = 2147483647.1

	for _ , kValue := range kk {
		var max float64
		for i:=0;i<len(xs);i++ {
			tmpXLen := xs[i] - xs[kValue-1]
			tmpYLen := ys[i] - ys[kValue-1]

			 // fmt.Println(tmpXLen, tmpYLen)

			l := math.Sqrt(float64((tmpXLen * tmpXLen) + (tmpYLen * tmpYLen)))
			// fmt.Println(kValue, l)
			if l > max {
				max = l
			}
		}

		if max < ans {
			ans = max
		}
		
	}

	fmt.Println(ans)
}
