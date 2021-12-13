package main

import (
	"fmt"
	"sort"
)

func main (){
	array := []int{-2, -1}
	result := SortedSquaredArray(array)
	fmt.Println(result)
}



func SortedSquaredArray(array []int) []int {
	for i,v := range array {
        array[i] = v * v
    }
	sort.Ints(array)
	return array
}


