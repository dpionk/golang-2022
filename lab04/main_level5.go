package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

type wynik struct {
	imie  string
	wynik int
	data  time.Time
}

func wylosuj() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

func dodaj_wynik(imie string, proby int, wyniki []wynik, gry int) []wynik {
	wynik_gracza := wynik{imie: imie, wynik: proby, data: time.Now()}
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

func wczytaj_wyniki() {

}

func zapisz_do_pliku(wyniki []wynik) {

	f, err := os.OpenFile("results.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	for _, el := range wyniki {
		if _, err = f.WriteString(el.imie + " " + strconv.Itoa(el.wynik) + " " + el.data.Format("2006-01-02 15:04:05") + "\n"); err != nil {
			panic(err)
		}
	}
	fmt.Println("Zapisano wyniki w pliku")
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
		zapisz_do_pliku(wyniki)
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
			zapisz_do_pliku(nowe_wyniki)
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
	fmt.Println("Witaj! Teraz będziesz zgadywać liczbę, którą wylosowałxm")
	fmt.Println("Aby wyjść z programu, wpisz 'koniec' ")
	game(wylosowana_liczba, 0, wyniki, 0)
}
