package main

import (
	"fmt"
	"strings"
)

func main() {
	var in string
	fmt.Scanf("%s", &in)

	fmt.Println(strings.Count(in, "1"))
}
