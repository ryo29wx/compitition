package main

import "fmt"

type listlist struct {
	slice []string
}

func main() {

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

}
