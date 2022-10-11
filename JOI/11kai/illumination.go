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
	field  [105][105]int
	dist   [105][105]bool
	height int
	width  int

	// 左上　右上
	sx = 0
	sy = 0
	gx = 104
	gy = 104
	// 奇数行
	d1x = [6]int{0, 1, 1, 1, 0, -1}
	d1y = [6]int{1, 1, 0, -1, -1, 0}
	// 偶数行
	d2x = [6]int{0, 1, 0, -1, -1, -1}
	d2y = [6]int{1, 0, -1, -1, 0, 1}
)

type point struct {
	x int
	y int
}

func main() {
	sc.Split(bufio.ScanWords)
	width = nextInt()
	height = nextInt()

	for i := 1; i <= height; i++ {
		for j := 1; j <= width; j++ {
			field[i][j] = nextInt()
		}
	}
	ans := bfs()
	fmt.Println(ans)

}

// とりあえず左上から
func bfs() int {
	counter := 0
	var q []point //point構造体型の配列
	q = append(q, point{sx, sy})

	for len(q) != 0 {
		current := q[0]
		x := current.x
		y := current.y
		q = q[1:]

		for d := 0; d < 6; d++ {
			var nextX int
			var nextY int
			if y%2 == 0 {
				nextX = x + d2x[d]
				nextY = y + d2y[d]
			} else {
				nextX = x + d1x[d]
				nextY = y + d1y[d]
			}

			if nextX < 0 || nextX > width+1 || nextY < 0 || nextY > height+1 {
				continue
			}
			if field[nextY][nextX] == 1 {
				counter++
				continue
			}

			if field[nextY][nextX] == 0 {
				if !dist[nextY][nextX] {
					dist[nextY][nextX] = true
					q = append(q, point{nextX, nextY})
				}
			}

			if nextX == gx && nextY == gy {
				break
			}
		}
	}
	return counter
}
