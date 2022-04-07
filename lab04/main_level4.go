package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type wynik struct {
	imie  string
	wynik int
}

func wylosuj() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

func dodaj_wynik(imie string, proby int, wyniki []wynik, gry int) []wynik {
	wynik_gracza := wynik{imie: imie, wynik: proby}
	wyniki = append(wyniki, wynik_gracza)
	return wyniki
}

func print_wyniki(wyniki []wynik) {

	sort.Slice(wyniki, func(i, j int) bool {
		return wyniki[i].wynik < wyniki[j].wynik
	})

	fmt.Println("Wyniki tej gry:")
	for _, el := range wyniki {
		fmt.Println(el.imie, el.wynik)
	}

}

func game(wylosowana_liczba int, proby int, wyniki []wynik, gry int) {
	fmt.Println("Podaj liczbę:")
	fmt.Println(wylosowana_liczba)
	var odpowiedz string
	proby++
	fmt.Scan(&odpowiedz)
	intVar, _ := strconv.Atoi(odpowiedz)
	switch {
	case odpowiedz == "koniec":
		fmt.Println("Żegnaj...")
		print_wyniki(wyniki)
	case intVar == wylosowana_liczba:
		fmt.Println("Brawo! Odgadłeś liczbę! Ilość prób: ", proby)
		fmt.Println("Podaj swoje imię")
		var imie string
		fmt.Scan(&imie)
		nowe_wyniki := dodaj_wynik(imie, proby, wyniki, gry)
		fmt.Println("Czy chcesz grać dalej? T/N")
		var kontynuuj string
		fmt.Scan(&kontynuuj)
		switch kontynuuj {
		case "T", "t":
			game(wylosuj(), 0, nowe_wyniki, gry+1)
		case "N", "n":
			fmt.Println("Żegnaj...")
			print_wyniki(nowe_wyniki)
		default:
			fmt.Println("Nieznana komenda :|")
		}
	case intVar < wylosowana_liczba:
		fmt.Println("Więcej...")
		game(wylosowana_liczba, proby, wyniki, gry)
	case intVar > wylosowana_liczba:
		fmt.Println("Mniej...")
		game(wylosowana_liczba, proby, wyniki, gry)
	}
}

func main() {
	var wyniki []wynik
	wylosowana_liczba := wylosuj()
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Witaj! Teraz będziesz zgadywać liczbę, którą wylosowałxm")
	fmt.Println("Aby wyjść z programu, wpisz 'koniec' ")
	game(wylosowana_liczba, 0, wyniki, 0)
}
