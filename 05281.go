package main

import (
	"fmt"
	//"strings"
)

func main() {
	var s string = "asenadcfish"
	m := make(map[string]int)
	//fmt.Println(len(s))
	for i := 0; i < len(s); i++ {
		//fmt.Println(s[i : i+1])
		m[s[i:i+1]]++
		fmt.Println(m)

	}
	fmt.Println(m)

}
