/*
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
	n := next()

	flg := false
	for i := 0; i < 10; i++ {
		rev := reverse(n)
		if n == rev {
			flg = true
			break
		}
		n = "0" + n

	}

	if flg {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
*/