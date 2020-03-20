package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Println("fib(", i, ") = ", Fib(i))
	}
}

func Fib(position int) int {
	if position <= 1 {
		return position
	}
	return fib(position-1) + fib(position-2)
}
