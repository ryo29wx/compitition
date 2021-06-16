package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	var test string
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%s", &test)

	fmt.Printf("%d %s\n", a+b+c, test)
}
