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
	s [][]rune
	t [][]rune
	N int
)

func main() {
	sc.Split(bufio.ScanWords)
	N = nextInt()

	for h := 0; h < N; h++ {
		str := next()
		s = append(s, []rune(str))
	}
	for h := 0; h < N; h++ {
		str := next()
		// t[h] = []rune(str)
		t = append(t, []rune(str))

	}

	if !count() {
		fmt.Println("No")
		return
	}

	flg := false
	for i := 0; i < 4; i++ {
		if isSame() {
			fmt.Println("Yes")
			flg = true
			break
		}

		if i != 3 {
			kaiten()
		}
	}

	if !flg {
		fmt.Println("No")
	}

}

func count() bool {
	sCnt := 0
	tCnt := 0
	for _, v := range s {
		for _, vv := range v {
			if vv == '#' {
				sCnt++
			}
		}
	}

	for _, v := range t {
		for _, vv := range v {
			if vv == '#' {
				tCnt++
			}
		}
	}

	if sCnt == tCnt {
		return true
	}
	return false
}

func kaiten() {
	l := len(s)
	for layer := 0; layer < l/2; layer++ {
		first := layer
		last := l - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first
			top := s[first][i]

			s[first][i] = s[last-offset][first]
			s[last-offset][first] = s[last][last-offset]
			s[last][last-offset] = s[i][last]
			s[i][last] = top
		}
	}
}

func findLeftTop(str string) (int, int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if str == "s" {
				if s[i][j] == '#' {
					return i, j
				}
			} else {
				if t[i][j] == '#' {
					return i, j
				}
			}
		}
	}
	return -1, -1
}

func isSame() bool {
	Si, Sj := findLeftTop("s")
	Ti, Tj := findLeftTop("t")
	offsetI := Ti - Si
	offsetJ := Tj - Sj
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			ii := i + offsetI
			jj := j + offsetJ
			if 0 <= ii && ii < N && 0 <= jj && jj < N {
				if s[i][j] != t[ii][jj] {
					return false
				}
			} else {
				if s[i][j] == '#' {
					return false
				}
			}
		}
	}
	return true
}
