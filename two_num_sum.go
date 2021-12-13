package main

import "fmt"

func main() {
	result := TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6},10)
	fmt.Println(result)
}


func TwoNumberSum(array []int, target int) []int {
	for i := 0; i < len(array)-1; i++ {
		primerValor := array[i]
		for j := i+1; j < len(array); j++ {
			segundoValor := array[j]
			if primerValor+segundoValor==target{
				return []int{primerValor,segundoValor}
			}
		}
	}
	return []int{}
}
