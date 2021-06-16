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
	N := nextInt()

	var s1 []string
	var s2 []string

	for i:=0; i<N; i++ {
		s := next()
		if strings.HasPrefix(s, "!") {
			s1 = append(s1, s)
		} else {
			s2 = append(s2, s)
		}
	}

	flg := false
	for i:=0; i<len(s2); i++ {
		comstr := "!" + s2[i]
		for _, val := range s1 {
			if comstr == val {
				fmt.Println(s2[i])
				flg = true
				break
			}
		}

		if flg {
			break
		}
	}

	if !flg {
		fmt.Println("satisfiable")
	}
}