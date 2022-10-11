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
	slice []int
	c     []int
)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	c = append(c, 0)
	for i := 1; i <= n; i++ {
		tmp := nextInt()
		slice = append(slice, tmp)
		c = append(c, c[i-1]+tmp)
	}

	for k := 1; k <= n; k++ {
		var ans int
		for j := 1; j <= n; j++ {
			if j-k < 0 {
				continue
			}

			ans1 := c[j] - c[j-k]
			if ans < ans1 {
				ans = ans1
			}
		}
		fmt.Println(ans)
	}

}
