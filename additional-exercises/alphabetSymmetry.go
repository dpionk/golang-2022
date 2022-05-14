// Consider the word "abode". We can see that the letter a is in position 1 and b is in position 2.
//  In the alphabet, a and b are also in positions 1 and 2. Notice also that d and e in abode occupy the positions
// they would occupy in the alphabet, which are positions 4 and 5.
// Given an array of words, return an array of the number of letters that occupy their positions in the alphabet
// for each word. For example,
// solve(["abode","ABc","xyzD"]) = [4, 3, 1]

package main

import (
	"fmt"
	"sort"
	"strings"
)

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "r", "s", "t", "u", "w", "x", "y", "z"}

func solve(slice []string) []int {
	var result []int
	for i := 0; i < len(slice); i++ {
		score := 0
		for j := 0; j < len(slice[i]); j++ {
			indexInAlphabet := sort.StringSlice(alphabet).Search(strings.ToLower(string(slice[i][j])))
			if indexInAlphabet == j {
				score++
			}
		}
		result = append(result, score)
	}
	return result
}

func main() {
	fmt.Println(solve([]string{"abode", "ABc", "xyzD"}))
}
