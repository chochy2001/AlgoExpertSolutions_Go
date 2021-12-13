package main

import "fmt"

func main() {
	//fn = nthFibonacci(n)
	// fn = fn-2 + fn-1

	value := 10
	fmt.Println(GetNthFib(value))

}

func GetNthFib(n int) int {
	// 0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987
	if n == 2 {
		return 1
	} else if n == 1 {
		return 0
	} else {
		return GetNthFib(n-1) + GetNthFib(n-2)
	}
}
