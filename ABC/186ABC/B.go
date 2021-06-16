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

	var x_s []float64
	var y_s []float64

	for i:=0; i<N; i++ {
		x := nextFloat()
		y := nextFloat()

		x_s = append(x_s, x)
		y_s = append(y_s, y)
	}


	var cnt int
	for len(x_s)>0 {
		x0 := x_s[0]
		y0 := y_s[0]

		for i:=1; i<len(x_s); i++ {
			if len(x_s) == 1 {
				break
			}
			xi := x_s[i]
			yi := y_s[i]

			var ans float64
			if x0 > xi {
				ans = (y0-yi)/(x0-xi)
			} else {
				ans = (yi-y0)/(xi-x0)
			}

			if ans >= -1 && ans <= 1 {
				cnt++
			}
		}
		x_s = unset(x_s)
		y_s = unset(y_s)
	}

	fmt.Println(cnt)
}

func unset(s []float64) []float64 {
	return append(s[:0], s[1:]...)
}