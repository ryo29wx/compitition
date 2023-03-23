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
	N := nextInt()

	for i := 0; i < (1 << N); i++ {
		candidate := ""
		for j := N - 1; j >= 0; j-- {
			// メモ : (i & (1 << j)) = 0 というのは、i の j ビット目（2^j の位）が 0 であるための条件。
			// 　　　頻出なので知っておくようにしましょう。
			if (i & (1 << j)) == 0 {
				candidate += "("
			} else {
				candidate += ")"
			}
		}
		I := hantei(candidate)
		if I == true {
			fmt.Println(candidate)
		}
	}
}

func hantei(S string) bool {
	dep := 0
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			dep += 1
		}
		if S[i] == ')' {
			dep -= 1
		}
		if dep < 0 {
			return false
		}
	}
	if dep == 0 {
		return true
	}
	return false
}
