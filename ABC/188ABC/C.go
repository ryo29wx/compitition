package main
 
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
 
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
	N := nextFloat()
	A := make([]int, 0)

	for i := 0; i < int(math.Pow(2.0, N)); i++ {
		A = append(A, nextInt())
	}

	B := &make([]int, 0)
	for i := 0; i < len(A); i=i+2 {
		if A[i] > a[i+1] {
			B  = append(B, A[i])
		} else {
			B  = append(B, A[i+1])
		}
	}

}

func memo() {
	B := make([]int, 0)
	for i := 0; i < len(A); i=i+2 {
		if A[i] > a[i+1] {
			B  = append(B, A[i])
		} else {
			B  = append(B, A[i+1])
		}
	}
}