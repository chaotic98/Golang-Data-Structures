package main

import "fmt"

func main() {
	fibonacci(7)
}

func fibonacci(n int) {
	fib := make([]int, n)

	fib[0], fib[1] = 0, 1

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	fmt.Println(fib)
}
