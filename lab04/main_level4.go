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

func game(wylosowana_liczba int, gry int) {
	gry++
	fmt.Println("Podaj swoje imię")
	var imie string
	fmt.Scan(&imie)
	fmt.Println("Podaj liczbę:")
	var odpowiedz string
	proby := 1
	fmt.Scan(&odpowiedz)
	intVar, _ := strconv.Atoi(odpowiedz)
	switch {
	case odpowiedz == "koniec":
		fmt.Println("Żegnaj...")
	case intVar == wylosowana_liczba:
		fmt.Println("Brawo! Odgadłeś liczbę!")
		fmt.Println("Czy chcesz grać dalej? T/N")
		var kontynuuj string
		fmt.Scan(&kontynuuj)
		switch kontynuuj {
		case "T", "t":
			game(wylosuj())
		case "N", "n":
			fmt.Println("Żegnaj...")
		default:
			fmt.Println("Nieznana komenda :|")
		}
	case intVar < wylosowana_liczba:
		fmt.Println("Więcej...")
		proby++
		game(wylosowana_liczba)
	case intVar > wylosowana_liczba:
		proby++
		fmt.Println("Mniej...")
		game(wylosowana_liczba)
	}
}

func main() {
	wyniki := map[string]map[string]int{}
	gry := 0
	wylosowana_liczba := wylosuj()
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Witaj! Teraz będziesz zgadywać liczbę, którą wylosowałxm")
	fmt.Println("Aby wyjść z programu, wpisz 'koniec' ")
	game(wylosowana_liczba)

}
