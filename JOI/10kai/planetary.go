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
	H     int
	W     int
	field [10][10]int
	sum   [10][10]int
)

func main() {
	sc.Split(bufio.ScanWords)
	H = nextInt()
	W = nextInt()

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			field[i][j] = nextInt()
		}
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			sum[i+1][j+1] = sum[i][j+1] + sum[i+1][j] - sum[i][j] + field[i][j]
		}
	}

	Q := nextInt()
	var ans []int
	for q := 0; q < Q; q++ {
		var x1, x2, y1, y2 int
		x1 = nextInt()
		x2 = nextInt()
		y1 = nextInt()
		y2 = nextInt()
		ans = append(ans, (sum[x2][y2] - sum[x1][y2] - sum[x2][y1] + sum[x1][y1]))
	}

	for _, v := range ans {
		fmt.Println(v)
	}

}
