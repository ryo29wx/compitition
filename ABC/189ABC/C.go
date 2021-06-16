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

	var slice []int
	for i:=0; i<N ; i++ {
		slice = append(slice, nextInt())
	}

	max := 0
	for i := 0; i <= N; i++ {
		for j := 0; j < i; j++ {
			newSlice := slice[j:i]
			fmt.Println(newSlice)
			min := 0
			for i, num := range newSlice {
				if i == 0 {
					min = num
					continue
				}

				if min >= num {
					min = num
				}
			}
			if max < (min*len(newSlice)) {
				max =  min * len(newSlice)
			}

		}
		
	}

	fmt.Println(max)

}
