package main


import (
	"fmt"
)


// お年玉が可能なのか判定する
func main() {
	var n int
	var y int
	fmt.Scanf("%v", &n)
	fmt.Scanf("%v", &y)

	flg := false

	for i:=0; i<=n; i++ {
		var sum int
		sum = sum + (10000 * i)

		for j:=0; j<=(n-i); j++ {
			sum = sum + (5000 * j)

			for k:=0; k<=(n-i-j); k++ {
				sum = sum + (1000 * k)
				fmt.Println(sum)
				if sum == y && (i+j+k) == n {
					fmt.Println(i)
					fmt.Println(j)
					fmt.Println(k)

					flg = true
					break
				}
			}
		}
	}

	if flg == false {
		fmt.Println("-1")
	}

}