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
	n := nextInt()
	s := nextInt()
	d := nextInt()

	flg := false
	for i := 0; i < n; i++ {
		x := nextInt()
		y := nextInt()

		if x < s {
			if y > d {
				flg = true
				break
			}
		}
	}

	if flg {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
