package main
 
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math"
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
	X := nextFloat()
	Y := nextFloat()

	if math.Abs(X - Y) < 3 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}