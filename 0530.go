package main

import "fmt"

func fibonacci() func() int {
	var n, m int = 0, 1
	return func() int {
		n, m = m, n+m
		return n
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
