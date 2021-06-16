package main

import (
	"bufio"
	"fmt"
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

func main() {
	sc.Split(bufio.ScanWords)
	N := nextInt()

	counter := make(map[int]int)
	for i := 0; i < N; i++ {
		An := nextInt()
		over := An % 200
		if val, ok := counter[over]; ok {
			counter[over] = val + 1
		} else {
			counter[over] = 1
		}
	}

	var result int
	for _, value := range counter {
		if value >= 2 {
			result = result + plus(value-1)
		}
	}

	fmt.Println(result)
}

func plus(x int) int {
	var count int
	for i := 1; i <= x; i++ {
		count = count + i
	}
	return count
}
