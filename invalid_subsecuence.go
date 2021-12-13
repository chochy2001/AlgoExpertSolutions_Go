package main

import "fmt"

func main(){
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence:= []int{1, 6, -1, 10}
	result := IsValidSubsequence(array, sequence)
	fmt.Printf("%v\n", result)
}


func IsValidSubsequence(array []int, sequence []int) bool {
	sequenceIndex := 0

	for _,v := range array {
		if sequence[sequenceIndex] == v {
			sequenceIndex++
		}
		if sequenceIndex == len(sequence) {
            return true
        }
	}
	return false
}

