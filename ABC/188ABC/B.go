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
	A := make([]int, N)
	B := make([]int, N)
	inner_pro := 0

	for i := 0; i < N; i++ {
		A[i] = nextInt()
	}

	for i := 0; i < N; i++ {
		B[i] = nextInt()
		inner_pro = inner_pro + (A[i] * B[i])
	}

	if inner_pro == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}