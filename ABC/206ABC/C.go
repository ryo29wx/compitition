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
	var slice []int
	for i := 0; i < n; i++ {
		slice = append(slice, nextInt())
	}
	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })

	cnt := 0
	ans := 1
	for i := 0; i < n-1; i++ {
		if slice[i] == slice[i++] {
			cnt++
			continue
		}

		for j := n-cnt; j > 0; j-- {
			ans = ans * j
		}
		
	}

}
