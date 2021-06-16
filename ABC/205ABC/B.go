/** package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
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
	var slice2 []int
	for i := 0; i < n; i++ {
		slice = append(slice, nextInt())
		slice2 = append(slice2, i+1)
	}

	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	//fmt.Println(slice)

	if reflect.DeepEqual(slice, slice2) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
