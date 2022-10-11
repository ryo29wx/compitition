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
	S [1010]string
	T [1010]string
)

func main() {
	sc.Split(bufio.ScanWords)
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		S[i] = next()
		T[i] = next()

	}

	flg := false
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if S[i] == S[j] {
				if T[i] == T[j] {
					flg = true
					break
				}
			}
		}
		if flg {
			break
		}
	}

	if flg {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
