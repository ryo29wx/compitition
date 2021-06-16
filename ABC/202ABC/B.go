package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

/**func main() {
	sc.Split(bufio.ScanWords)
	var S string
	fmt.Scan(&S)

	//re := reverse(S)
	var ans string
	for _, v := range S {
		switch v {
		case '6':
			ans = "9" + ans
		case '9':
			ans = "6" + ans
		default:
			ans = string(v) + ans
		}
	}
	fmt.Println(ans)

}
**/
func main() {
	sc.Split(bufio.ScanWords)
	var S string
	fmt.Scan(&S)

	re := reverse(S)
	re = strings.Replace(re, "6", "r", -1)
	re = strings.Replace(re, "9", "6", -1)
	fmt.Println(strings.Replace(re, "r", "9", -1))

}

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
