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
	m := nextInt()
	t := nextInt()

	remain := n
	beforeB := 0
	for i:=0; i<m; i++ {
		a := nextInt()
		b := nextInt()

		remain = remain - (a - beforeB)
		if remain <= 0 {
			break
		}

		remain = remain + (b - a)
		beforeB = b

		if remain >= n {
			remain = n
		}
	}
	if remain <= 0 {
		fmt.Printf("No\n")
	} else {
		remain = remain - (t - beforeB)
		if remain <= 0 {
			fmt.Printf("No\n")
		} else {
			fmt.Printf("Yes\n")
		}
	}

}