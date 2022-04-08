package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
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

func wczytaj_wyniki() map[string]int64 {
	file, err := os.Open("results.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wyniki := make(map[string]int64)
	for scanner.Scan() {
		v := strings.Split(scanner.Text(), " ")
		intVar, _ := strconv.ParseInt(v[1], 0, 16)
		if _, found := wyniki[v[0]]; found {
			if wyniki[v[0]] > intVar {
				wyniki[v[0]] = intVar
			}
		} else {
			wyniki[v[0]] = intVar
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return wyniki
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

func sprawdz_rekord(imie string, wynik int64, poprzednie_wyniki map[string]int64) {
	if _, found := poprzednie_wyniki[imie]; found {
		if poprzednie_wyniki[imie] > wynik {
			fmt.Println("Brawo! Pobito własny rekord! Wcześniejszy: " + strconv.FormatInt(poprzednie_wyniki[imie], 10) + " Aktualny: " + strconv.FormatInt(wynik, 10))
		}
	}
}

func game(wylosowana_liczba int, proby int, wyniki []wynik, gry int, wyniki_z_pliku map[string]int64) {
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
		sprawdz_rekord(imie, int64(proby), wyniki_z_pliku)
		nowe_wyniki := dodaj_wynik(imie, proby, wyniki, gry)
		fmt.Println("Czy chcesz grać dalej? T/N")
		var kontynuuj string
		fmt.Scan(&kontynuuj)
		switch kontynuuj {
		case "T", "t":
			game(wylosuj(), 0, nowe_wyniki, gry+1, wyniki_z_pliku)
		case "N", "n":
			fmt.Println("Żegnaj...")
			print_wyniki(nowe_wyniki)
			zapisz_do_pliku(nowe_wyniki)
		default:
			fmt.Println("Nieznana komenda :|")
		}
	case intVar < wylosowana_liczba:
		fmt.Println("Więcej...")
		game(wylosowana_liczba, proby, wyniki, gry, wyniki_z_pliku)
	case intVar > wylosowana_liczba:
		fmt.Println("Mniej...")
		game(wylosowana_liczba, proby, wyniki, gry, wyniki_z_pliku)
	}
}

func main() {
	var wyniki []wynik
	wylosowana_liczba := wylosuj()
	wyniki_z_pliku := wczytaj_wyniki()
	fmt.Println("Witaj! Teraz będziesz zgadywać liczbę, którą wylosowałxm")
	fmt.Println("Aby wyjść z programu, wpisz 'koniec' ")
	game(wylosowana_liczba, 0, wyniki, 0, wyniki_z_pliku)
}
