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
	A := nextInt()
	B := nextInt()
	C := nextInt()
	X := nextInt()
	Y := nextInt()

	ans := A*X + B*Y
	var large int
	if X > Y {
		large = X * 2
	} else {
		large = Y * 2
	}

	// ABを積んでいく
	for i := 0; i <= large; i = i + 2 {
		abCost := C * i
		aCount := X - i/2
		bCount := Y - i/2
		if aCount < 0 {
			aCount = 0
		}
		if bCount < 0 {
			bCount = 0
		}

		sumCost := A*aCount + B*bCount + abCost
		if sumCost < ans {
			ans = sumCost
		}
	}

	fmt.Println(ans)

}
