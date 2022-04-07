package main

import (
	"fmt"
	"strings"
)

var lista_slow_dla_dnia_urodzin = map[int]string{1: "wspaniałomyślny", 2: "kreatywny", 3: "przepiękny",
	4: "interesowny", 5: "bezinteresowny", 6: "wielki", 7: "katastroficzny", 8: "malutki",
	9: "słodki", 10: "jakiś", 11: "zabawny", 12: "wkurzający", 13: "wkurzony", 14: "nieśmieszny",
	15: "cringowy", 16: "zadowolony", 17: "szczęśliwy", 18: "gotowy do działania", 19: "fałszywy", 20: "wierny",
	21: "niewierny", 22: "niebezpieczny", 23: "aspołeczny", 24: "gładki", 25: "niewidomy",
	26: "wąski", 27: "wygodny", 28: "zachwycony", 29: "bezrobotny", 30: "miły", 31: "strachliwy"}

var lista_slow_dla_imion = map[string]string{"A": "banan", "B": "stół", "C": "komputer", "D": "piesek",
	"E": "spinacz", "F": "szczur", "G": "kabel", "H": "ogór", "I": "kotek", "J": "robal",
	"K": "hipster", "L": "piłkarz", "M": "tukan", "N": "gepard", "O": "wróbelek", "P": "kubek",
	"R": "Rafalisko", "S": "łoś", "T": "orangutan", "U": "ciastek", "W": "pomylątko",
	"X": "xenon", "Y": "yeti", "Z": "but"}

var lista_slow_dla_nazwisk = map[string]string{"A": "płacze", "B": "spadł z rowerka", "C": "udławił się ciastem", "D": "tańczy na parkiecie",
	"E": "pragnie zjeść pomelo", "F": "koduje program w go", "G": "rozlał mleko", "H": "zepsuł komputer", "I": "siedzi na fotelu", "J": "skacze wysoko",
	"K": "kłoci się z mamą", "L": "wyjeżdża na wakacje", "M": "sprząta pokój", "N": "leży w szpitalu", "O": "wagaruje", "P": "gra w lola",
	"R": "buduje dom w minecraft", "S": "śpiewa do mikrofonu", "T": "scamuje ludzi w internecie", "U": "robi skany swojej ręki", "W": "myje zęby",
	"X": "pisze autobiografię", "Y": "chodzi na politechnikę", "Z": "kocha ug całym sercem"}

func ksywaMaker(d int, i string, n string) string {
	return lista_slow_dla_dnia_urodzin[d] + " " + lista_slow_dla_imion[i] + " " + lista_slow_dla_nazwisk[i]
}

func ksywa() {
	var dzien int
	fmt.Println("Proszę podać dzień urodzenia (liczba): ")
	fmt.Scan(&dzien)
	fmt.Println("Podano: ", dzien)
	var pierwsza_litera_imienia string
	fmt.Println("Proszę podać pierwszą literę imienia: ")
	fmt.Scan(&pierwsza_litera_imienia)
	var pierwsza_litera_nazwiska string
	fmt.Println("Proszę podać pierwszą literę nazwiska: ")
	fmt.Scan(&pierwsza_litera_nazwiska)
	ksywa := ksywaMaker(dzien, strings.ToUpper(pierwsza_litera_imienia), strings.ToUpper(pierwsza_litera_nazwiska))
	fmt.Println("Twoja super ksywa to:")
	fmt.Println(ksywa)
}

func main() {
	ksywa()
}
