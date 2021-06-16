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
	var s string
	fmt.Scan(&s)

	var ans int
	for i := 0; i <= 9999; i++ {
		var memo [10]bool
		now := i
		for j := 0; j < 4; j++ {
			memo[now%10] = true
			now /= 10
		}

		ansFlg := true
		for j, s := range s {
			if s == 'o' && !memo[j] {
				ansFlg = false
			}
			if s == 'x' && memo[j] {
				ansFlg = false
			}
		}
		if ansFlg {
			ans++
		}
	}
	fmt.Println(ans)

}
