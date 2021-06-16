package main 


import (
	"fmt"
)

func main() {

	var n, m, t int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &m)
	fmt.Scanf("%d", &t)


	inMap := make(map[int]int)
	for i:=0; i<=m; i++ {
		var a, b int
		fmt.Scanf("%d", &a)
		fmt.Scanf("%d", &b)

		inMap[a] = b
	}

	remainQuant := n
	// flg := 0
	// beforeK := 0
	beforeV := 0
	for k,v := range inMap {
		remainQuant = remainQuant - (k - beforeV) 
		fmt.Printf("%d\n", remainQuant)
		if remainQuant <= 0 {
			break
		}
		remainQuant = remainQuant + (v - k)
		if remainQuant >= n {
			remainQuant = n
		}
		fmt.Printf("%d\n", remainQuant)
		beforeV = v

	}
	remainQuant = remainQuant - (t - beforeV)
	fmt.Printf("%d\n", remainQuant)

	if remainQuant <= 0 {
		fmt.Printf("No\n")
	} else {
		fmt.Printf("Yes\n")
	}

}