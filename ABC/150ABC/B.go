package main

import (
	"bufio"
	"fmt"
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

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	S := next()

	fmt.Printf("%v", n)

	r := strings.Count(S, "ABC")
	fmt.Printf("%v", r)

}

