package main 


import (
	"fmt"
)

func main() {

	var a, b, c, d int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &d)


	var arr[4] int = [4]int{a,b,c,d}
	var min int = a
	for _,v := range arr {
		if v <= min {
			min = v
		}
	}

	fmt.Printf("%d\n", min)
}