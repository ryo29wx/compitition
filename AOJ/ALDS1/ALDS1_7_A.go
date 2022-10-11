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


type Node struct {
	parent int
	left int //child
	right int //sibling
}
 
func (node *Node) print(i int, depth *[]int, nodes *[]Node) {
	p := (*node).parent
	d := (*depth)[i]
	
	kind := ""
	if (*node).parent == -1 {
		kind = "root"
	} else if (*node).left == -1 {
		kind = "leaf"
	} else {
		kind = "internal node"
	}
	
	var children []int
	if (*node).left != -1 {
		children = append(children, (*node).left)
		var pre Node
		pre = (*nodes)[(*node).left]
		for pre.right != -1 {
			children = append(children, pre.right)
			pre = (*nodes)[pre.right]
		}
		
	}
	
	fmt.Printf("node %v: parent = %v, depth = %v, %v, %v \n", i, p, d, kind, children)
}

// 
func setDepth(u, p int, depth *[]int, nodes *[]Node) {
	(*depth)[u] = p
	if (*nodes)[u].right != -1 {
		setDepth(((*nodes)[u]).right, p, depth, nodes)
	}
	
	if (*nodes)[u].left != -1 {
		setDepth(((*nodes)[u]).left, p+1, depth, nodes)
	}
	
}

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()
	tmpN := make([]Node, n)
	tmpD := make([]int, n)

	nodes := &tmpN
	depth := &tmpD
	
	for i:=0;i<n;i++ {
		(*nodes)[i].parent = -1
		(*nodes)[i].left = -1
		(*nodes)[i].right = -1
	}
	
	for i:=0;i<n;i++ {
		v := nextInt()
		nOfChildren := nextInt()
		pre := 0
		fmt.Println(v, nOfChildren)
		for j:=0;j<nOfChildren;j++ {
			c := nextInt()
			if j == 0 {
				(*nodes)[v].left = c
			} else {
				(*nodes)[pre].right = c
			}
			pre = c
			(*nodes)[c].parent = v
			
		}
		
	}
	fmt.Println(*nodes)
	setDepth(0, 0, depth, nodes)

	for k, v := range (*nodes) {
		v.print(k, depth, nodes)
	}
}