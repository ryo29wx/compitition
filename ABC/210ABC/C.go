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
	n := nextInt()
	k := nextInt()

	var slice []int
	for i := 0; i < n; i++ {
		slice = append(slice, nextInt())
	}
	mp := make(map[int]int)
	for i := 0; i < k; i++ {
		mp[slice[i]]++
	}

	ans := len(mp)
	for i := k + 1; i <= n; i++ {
		mp[slice[i-1]]++
		mp[slice[i-k-1]]--
		if mp[slice[i-k-1]] == 0 {
			delete(mp, slice[i-k-1])
		}
		if len(mp) > ans {
			ans = len(mp)
		}
	}

	fmt.Println(ans)

}