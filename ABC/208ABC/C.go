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
	k := nextInt()

	var slice []int
	var slice2 []int
	for i := 0; i < n; i++ {
		tmp := nextInt()
		slice = append(slice, tmp)
		slice2 = append(slice2, tmp)
	}

	ans1 := k / n
	remaind := k % n

	sort.SliceStable(slice, func(i, j int) bool { return slice[i] < slice[j] })

	for _, v := range slice2 {

		if remaind != 0 && v <= slice[remaind-1] {
			fmt.Println(ans1 + 1)
		} else {
			fmt.Println(ans1)
		}
	}

}
