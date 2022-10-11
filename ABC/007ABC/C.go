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
	// 頂点の移動を表す
	dx = [4]int{1, 0, -1, 0}
	dy = [4]int{0, 1, 0, -1}
)

type point struct {
	x int
	y int
}

func main() {
	sc.Split(bufio.ScanWords)

	////////////////////////
	// 入力
	////////////////////////
	height := nextInt()
	width := nextInt()

	sy := nextInt() - 1
	sx := nextInt() - 1
	gy := nextInt() - 1
	gx := nextInt() - 1

	var field [][]rune
	for i := 0; i < height; i++ {
		field = append(field, []rune(next()))
	}

	///////////////////////
	// BFSの初期設定
	// 1. まだ探索していない事を表現する配列を用意
	// 2. これから訪問予定の座標をいれるキュー
	///////////////////////

	// 未訪問の座標は-1
	var dist [50][50]int
	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			dist[i][j] = -1
		}
	}

	dist[sx][sy] = 0

	var q []point //point構造体型の配列
	q = append(q, point{sx, sy})

	for len(q) != 0 {
		current := q[0]
		x := current.x
		y := current.y
		q = q[1:]

		for d := 0; d < 4; d++ {
			nextX := x + dx[d]
			nextY := y + dy[d]

			if nextX < 0 || nextX >= height || nextY < 0 || nextY >= width {
				continue
			}
			if string(field[nextX][nextY]) == "#" {
				continue
			}

			if dist[nextX][nextY] == -1 {
				q = append(q, point{nextX, nextY})
				dist[nextX][nextY] = dist[x][y] + 1
				// fmt.Println(nextX, nextY, dist[nextX][nextY])
				if nextX == gy && nextY == gx {
					fmt.Println(dist[nextX][nextY])
				}
			}
		}

	}

}
