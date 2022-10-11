/** package main

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
	var n int
	fmt.Scan(&n)
	set := make(map[string]struct{}, 1000)
	for i := 0; i < n; i++ {
		s := next()
		t := next()
		st := s + t
		set[st] = struct{}{}
	}

	if len(set) != n {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
