package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"sort"
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

	var mass [][]int
	maxVal := 0
	maxX := 0
	maxY := 0
	for i:=0;i<n;i++ {
		var low []int
		for j:=0;j<n;j++ {
			fmt.Println("fuga" , j)
			tmp := nextInt()
			fmt.Println("fuga" , tmp)

			if tmp > maxVal {
				maxVal = tmp
				maxX = j
				maxY = i
			}
			low = append(low, tmp)
		}
		mass = append(mass, low)
		fmt.Println("hoge" , i)
	}

	fmt.Println(u1(&mass, maxX, maxY, n))

}


func u1(mass *[][]int, maxX, maxY int, n int) int {
	xx := []int{-1, 0, 1}
	yy := []int{-1, 0, 1}
	
	var diff []int
	tmpA := strconv.Itoa((*mass)[maxX][maxY])
	for nextX := range xx {
		for nextY := range yy {
			if nextX == 0 && nextY == 0 {
				continue
			}
			nX := maxX
			nY := maxY
			for i:=0;i<n;i++ {
				nX = nX + nextX
				nY = nY + nextY
				tmpA = tmpA + strconv.Itoa((*mass)[nX][nY])
			}

			h, _ := strconv.Atoi(tmpA)
			diff = append(diff, h)

		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(diff)))
	return diff[0]
}
