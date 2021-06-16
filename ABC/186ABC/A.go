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
	a := nextInt()
	b := nextInt()

	x1 := a/100 
	tmp1 := a%100
	y1 := tmp1/10
	tmp1 = tmp1%10
	z1 := tmp1

	x2 := b/100 
	tmp2 := b%100
	y2 := tmp2/10
	tmp2 = tmp2%10
	z2 := tmp2

	ans1 := x1+y1+z1
	ans2 := x2+y2+z2

	if ans1 > ans2 {
		fmt.Println(ans1)
	} else {
		fmt.Println(ans2)
	}

}