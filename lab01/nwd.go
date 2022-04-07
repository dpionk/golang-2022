package main

import "fmt"

func nwd(a int, b int) int {
	if a == b {
		return a
	} else if a > b {
		return nwd(a-b, b)
	} else {
		return nwd(a, b-a)
	}
}

func main() {
	fmt.Println(nwd(24, 32))
}
