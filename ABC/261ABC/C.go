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
	q := nextInt()

	st := next()
	
	rr := []rune(st)
	for i:=0;i<q;i++ {
		t := nextInt()
		x := nextInt()
		if t == 1 {
			r := rr[x:n-1]						
			reverse(&r)
			rr = append(r, rr[0:]...)
			rr = rr[:n]
			fmt.Println(string(rr))
		} else {
			fmt.Println(string(rr[x-1]))
		}
	}

}

func reverse(l *[]rune) {
	for i := 0; i < len(*l) / 2; i++ {
    	(*l)[i], (*l)[len(*l) - i - 1] = (*l)[len(*l) - i - 1], (*l)[i]
	}
}