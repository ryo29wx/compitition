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

	ans := 0
	snapshotBit := 0b00000000
	for i:=0;i<n;i++ {
		initBit := 0b00000001
		
		t := snapshotBit | initBit

		left := t << nextInt()
		
		count := left >> 4

		if count == 15 {
			ans = ans + 4
		} else if count == 14 || count == 13 || count == 11 || count == 7 {
			ans = ans + 3
		} else if count == 12 || count == 10 || count == 9 || count == 6 || count == 5 || count == 3 {
			ans = ans + 2
		} else if count == 8 || count == 4 || count == 2 || count == 1 {
			ans = ans + 1
		} else {

		}
		snapshotBit = left & 0b00001111
		// fmt.Printf("%b \n", snapshotBit)

	}
	fmt.Println(ans)


}
