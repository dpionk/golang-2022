package main

import (
	"fmt"
	"math/rand"
	"time"
)

func game(wylosowana_liczba int) {
	fmt.Println("Podaj liczbę:")
	var zgadywana_liczba int
	fmt.Scan(&zgadywana_liczba)
	switch {
	case zgadywana_liczba == wylosowana_liczba:
		fmt.Println("Brawo! Odgadłeś liczbę!")
	case zgadywana_liczba < wylosowana_liczba:
		fmt.Println("Więcej...")
		game(wylosowana_liczba)
	case zgadywana_liczba > wylosowana_liczba:
		fmt.Println("Mniej...")
		game(wylosowana_liczba)
	}
}

func main() {
	wylosowana_liczba := rand.Intn(100)
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Witaj! Teraz będziesz zgadywać liczbę, którą wylosowałxm")
	game(wylosowana_liczba)

}
