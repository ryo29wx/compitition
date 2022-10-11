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
	field  [][]rune
	height int
	width  int
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
	height = nextInt()
	width = nextInt()
	chees := nextInt()

	for i := 0; i < height; i++ {
		field = append(field, []rune(next()))
	}

	ans := 0
	for i := 0; i < chees; i++ {
		tmp := bfs(i)
		ans += tmp
		// ans += bfs(i)

	}
	fmt.Println(ans)

}

func bfs(s int) int {
	// 都度初期化する
	var dist [1005][1005]int
	var sx, sy int
	var gx, gy int
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			dist[i][j] = -1
			if s == 0 {
				if field[i][j] == 'S' {
					sx = j
					sy = i
					dist[i][j] = 0
				}
				if field[i][j] == '1' {
					gx = j
					gy = i
				}
			} else {
				if field[i][j] == rune(s)+48 {
					sx = j
					sy = i
					dist[i][j] = 0
				}
				if field[i][j] == rune(s+1)+48 {
					gx = j
					gy = i
				}
			}
		}
	}
	//fmt.Println("start", sx, sy)

	var q []point //point構造体型の配列
	q = append(q, point{sx, sy})

	count := 0
	for len(q) != 0 {
		current := q[0]
		x := current.x
		y := current.y
		q = q[1:]

		for d := 0; d < 4; d++ {
			nextX := x + dx[d]
			nextY := y + dy[d]

			if nextX < 0 || nextX >= width || nextY < 0 || nextY >= height {
				continue
			}
			if string(field[nextY][nextX]) == "X" {
				continue
			}

			if dist[nextY][nextX] == -1 {
				q = append(q, point{nextX, nextY})
				dist[nextY][nextX] = dist[y][x] + 1
				if nextX == gx && nextY == gy {
					return dist[nextY][nextX]
				}
			}
		}
		count++

	}
	return 0
}
