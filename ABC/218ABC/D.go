/** package main

import (
	"bufio"
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

type point struct {
	x int
	y int
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	var p [2010]point
	for i := 0; i < n; i++ {
		bob := point{nextInt(), nextInt()}
		p[i] = bob
	}

	var kx []point
	var ky []point

	for i := 0; i < len(p); i++ {
		for j := i + 1; j < len(p); j++ {
			sX := p[i].x
			sY := p[i].y
			if p[j].x == sX {
				ky = append(ky, p[j])
			} else if p[j].y == sY {
				kx = append(kx, p[j])
			}

		}
	}

}
