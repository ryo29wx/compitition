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
	S := next()
	Q := nextInt()

	for i := 0; i < Q; i++ {
		T := nextInt()
		A := nextInt()
		B := nextInt()

		if T == 1 {
			a := string(getRuneAt(S, A-1))
			b := string(getRuneAt(S, B-1))
			fore := S[:A-1]
			intermediate := S[A : B-1]
			latter := S[B:]

			S = fore + b + intermediate + a + latter

		} else {
			fore := S[:N]
			latter := S[N:]
			S = latter + fore

		}

	}
	fmt.Println(S)

}

func getRuneAt(s string, i int) rune {
	rs := []rune(s)
	return rs[i]
}
