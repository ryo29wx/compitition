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

	mapEx := make(map[int]int)

	for i := 0; i < n; i++ {
		a := nextInt()
		b := nextInt()
		if key, val := mapEx[a]; val {
			mapEx[a] = key + b
		} else {
			mapEx[a] = b
		}
	}
	//fmt.Println("====map====")
	//fmt.Println(mapEx)

	keys := make([]int, 0, len(mapEx))
	for key := range mapEx {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	//fmt.Println("====key====")
	//fmt.Println(keys)

	for _, key := range keys {
		if k-key < 0 {
			// fmt.Println("break!")
			break
		} else {
			k = k + mapEx[key]
		}
	}

	fmt.Println(k)

}
