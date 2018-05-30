package main

import (
	//"code.google.com/p/go-tour/wc"
	"github.com/moby/moby.git"
)

func WordCount(s string) map[string]int {
	m = make(map[string]int)
	for i := 0; i < len(s); i++ {
		m[s[i:s+1]]++
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
