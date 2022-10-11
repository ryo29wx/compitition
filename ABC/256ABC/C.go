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
	h1 := nextInt()
	h2 := nextInt()
	h3 := nextInt()
	w1 := nextInt()
	w2 := nextInt()
	w3 := nextInt()

	ans := 0
	for m11:=1;m11<=h1-2;m11++ {
		for m12:=1;m12<=h1-2;m12++ {
			m13 := (h1 - m11 - m12)
			if m13 <= 0 {
				continue
			}
			for m21:=1;m21<=h2-2;m21++ {
				for m22:=1;m22<=h2-2;m22++ {
					m23 := (h2 - m21 - m22) 
					if m23 <= 0 {
						continue
					}
					for m31:=1;m31<=h3-2;m31++ {
						for m32:=1;m32<=h3-2;m32++ {
							m33 := (h3 - m31 - m32) 
							if m33 <= 0 {
								continue
							}
							
							if (m11+m21+m31) == w1 && (m12+m22+m32) == w2 && (m13+m23+m33) == w3  {
								ans++
							}
							
						}
					}
					
				}
			}
		}
	}
	fmt.Println(ans)
}
