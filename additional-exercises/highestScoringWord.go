// Given a string of words, you need to find the highest scoring word.
// Each letter of a word scores points according to its position in the alphabet: a = 1, b = 2, c = 3 etc.
// You need to return the highest scoring word as a string.
// If two words score the same, return the word that appears earliest in the original string.
// All letters will be lowercase and all inputs will be valid.

package main

import (
	"fmt"
	"strings"
)

type highestScoring struct {
	word  string
	score int
}

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "r", "s", "t", "u", "w", "x", "y", "z"}

func findLetterScore(letter string) int {
	for i := 0; i < len(alphabet); i++ {
		if alphabet[i] == letter {
			return i + 1
		}
	}
	return 0
}
func highestScoringWord(s string) string {
	splitedString := strings.Split(s, " ")
	highestScoringWordValue := highestScoring{"", 0}
	for i := 0; i < len(splitedString); i++ {
		score := 0
		for j := 0; j < len(splitedString[i]); j++ {
			score += findLetterScore(string(splitedString[i][j]))
		}
		if score > highestScoringWordValue.score {
			highestScoringWordValue.score = score
			highestScoringWordValue.word = splitedString[i]
		}
	}
	return highestScoringWordValue.word
}

func main() {
	fmt.Println(highestScoringWord("what time are we climbing up the volcano"))
	fmt.Println(highestScoringWord("man i need a taxi up to ubud"))
	fmt.Println(highestScoringWord("take me to semynak"))
	fmt.Println(highestScoringWord("aa b"))
	fmt.Println(highestScoringWord("aaa b"))
}
