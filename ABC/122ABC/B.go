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
	in := next()

	var max, ans int

	for i := 0; i < len(in); i++ {
		char := in[i]
		if char == 'A' || char == 'G' || char == 'T' || char == 'C' {
			max++
			if ans < max {
				ans = max
			}
		} else {
			max = 0
		}
	}

	fmt.Println(ans)

}
