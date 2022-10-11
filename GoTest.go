package main

import (
	"fmt"
	"sort"
)

type listlist struct {
	slice []string
}

type testlist struct {
	A int
}

func main() {
	testRune := []rune("abcd")
	fmt.Println(testRune[0] - 'a')
	fmt.Println('a')

	// slice 初期化
	var slice []string

	slice = append(slice, "aaaa")
	slice = append(slice, "bbbb")

	fmt.Println(slice)

	// sliceの中にslice
	var sliceSlice [][]string

	sliceSlice = append(sliceSlice, slice)
	sliceSlice = append(sliceSlice, slice)

	fmt.Println(sliceSlice)

	// map 初期化
	mapEx := make(map[int]int)
	mapEx[1] = 2
	fmt.Println(mapEx)

	// mapのvalueがslice
	mapSlice := map[int][]string{}
	mapSlice[1] = slice
	fmt.Println(mapSlice)

	// mapのvalueがmap
	mapMap := map[int]map[int]int{}
	mapMap[1] = mapEx
	fmt.Println(mapMap)

	// sliceの中にmap
	// var sliceMap []map[int]int
	sliceMap := []map[int]int{}
	sliceMap = append(sliceMap, mapEx)
	sliceMap = append(sliceMap, mapEx)

	fmt.Println(sliceMap)

	testInt := 10
	fmt.Println(testInt)

	pointer(&testInt)
	fmt.Println(testInt)

	//BIT_FLAG_0 := 1 << 0
	//BIT_FLAG_1 := 1 << 1
	//BIT_FLAG_2 := 1 << 2
	//BIT_FLAG_3 := 1 << 3
	//BIT_FLAG_4 := 1 << 4
	//BIT_FLAG_5 := 1 << 5
	//BIT_FLAG_6 := 1 << 6
	//BIT_FLAG_7 := 1 << 7
	//bit := BIT_FLAG_1 | BIT_FLAG_3 | BIT_FLAG_5
	//fmt.Println(bit & BIT_FLAG_3)

	a := 45 // 101101
	b := 25 // 011001
	fmt.Println(a)
	fmt.Println(b &^ a)
	fmt.Println(fmt.Sprintf("%b", b&^a))

	n := 5
	k := 3

	bit := (1 << k) - 1 // bit = {0, 1, 2}
	for ; bit < (1 << n); bit = nextCombination(bit) {
		var s []int
		for i := 0; i < n; i++ {
			if 1<<i == bit&(1<<i) {
				s = append(s, i)
			}
		}

		fmt.Printf("%v : {", bit)
		for _, v := range s {
			fmt.Printf("%v,", v)
		}
		fmt.Printf("}\n")
	}

	//x := []int{0, 1, 1, 3}
	//fmt.Println(0, x)

	//y := []string{"aaa", "bbb", "ccc", "bca"}
	//fmt.Println(0, y)

	r := []rune("abc")
	fmt.Println(0, string(r))

	var intSlice []int
	for _, v := range r {
		intSlice = append(intSlice, int(v))
	}
	for i := 1; nextPermutation(sort.IntSlice(intSlice)); i++ {
		var rSlice []rune
		for _, v := range intSlice {
			rSlice = append(rSlice, rune(v))
		}
		fmt.Println(i, string(rSlice))
	}
}

func nextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func nextCombination(sub int) int {
	x := sub & -sub
	y := sub + x
	return (((sub & ^y) / x) >> 1) | y
}

func pointer(test *int) {
	*test++
	fmt.Println(*test)
}
