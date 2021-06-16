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
	q := nextInt()
	for x := 0; x <= 100; x++ {
		var flg1 bool
		for y := 0; y <= 100; y++ {
			var tmp int
			var flg bool
			for k, v := range slice {
				xi := v[0]
				yi := v[1]
				hi := v[2]
				h := float64(hi) + math.Abs(float64(xi-x)) + math.Abs(float64(yi-y))
				if h < 0 {
					h = 0
				}
				if k == 0 {
					tmp = int(h)
				} else {
					if tmp == int(h) {
						flg = true
						continue
					} else {
						flg = false
						break
					}
				}
			}

			if flg {
				flg1 = true
				fmt.Println(x, y, tmp)
				break
			}
		}

		if flg1 {
			break
		}
	}
}
