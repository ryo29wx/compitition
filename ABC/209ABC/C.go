package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	n := nextInt()
	var c []int
	for i := 0; i < n; i++ {
		c = append(c, nextInt())
	}
	sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })

	ans := 1
	for i := 0; i < len(c); i++ {
		if c[i]-i <= 0 {
			ans = 0
		}
		ans = ans * (c[i] - i) % 1000000007
	}

	fmt.Println(ans)
}
