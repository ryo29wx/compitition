package main


import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%v", &n)

	var d_slice []int
	for i := 0; i<n; i++ {
		var d int
		fmt.Scanf("%v", &d)
		d_slice = append(d_slice, d)
	}

	set := make(map[int]struct{})
	for _, d := range d_slice {
		set[d] = struct{}{}
	}

	fmt.Println(len(set))
}