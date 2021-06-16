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
	m := nextInt()

	var slice []map[int]int{}
	for i := 0; i < m; i++ {
		tmpMap := map[int]int{}
		k := nextInt()
		v := nextInt()
		map[k] = v
		slice = append(slice, tmpMap)
	}

	for i := 1; i <= 2000; i++ {
		for j := 1; j <= 2000; j++ {
			
		}
	}

	fmt.Println(ans)

}
