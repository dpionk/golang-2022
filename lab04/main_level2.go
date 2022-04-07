package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func wylosuj() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

func game(wylosowana_liczba int) {
	fmt.Println("Podaj liczbę:")
	var odpowiedz string
	fmt.Scan(&odpowiedz)
	intVar, _ := strconv.Atoi(odpowiedz)
	switch {
	case odpowiedz == "koniec":
		fmt.Println("Żegnaj...")
	case intVar == wylosowana_liczba:
		fmt.Println("Brawo! Odgadłeś liczbę!")
	case intVar < wylosowana_liczba:
		fmt.Println("Więcej...")
		game(wylosowana_liczba)
	case intVar > wylosowana_liczba:
		fmt.Println("Mniej...")
		game(wylosowana_liczba)
	}
}

func main() {
	wylosowana_liczba := wylosuj()
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Witaj! Teraz będziesz zgadywać liczbę, którą wylosowałxm")
	fmt.Println("Aby wyjść z programu, wpisz 'koniec' ")
	game(wylosowana_liczba)

}
