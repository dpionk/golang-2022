//przykładowe wywołanie: go run program.go -liczba1 20 -liczba2 5

package main

import (
	"flag"
	"fmt"
)

func main() {
	liczba1 := flag.Int("liczba1", -10000, "pierwsza liczba do sumy")
	liczba2 := flag.Int("liczba2", -10000, "druga liczba do sumy")
	flag.Parse()
	if *liczba1 == -10000 {
		fmt.Print("Proszę podać pierwszą liczbę: ")
		var n int
		fmt.Scan(&n)
		fmt.Println("Suma podanych liczb", n, "i", *liczba2, "jest równa", n+*liczba2)
	} else if *liczba2 == -10000 {
		fmt.Print("Proszę podać drugą liczbę: ")
		var n int
		fmt.Scan(&n)
		fmt.Println("Suma podanych liczb", *liczba1, "i", n, "jest równa", *liczba1+n)
	} else {
		fmt.Println("Suma podanych liczb", *liczba1, "i", *liczba2, "jest równa", *liczba1+*liczba2)
	}
}
