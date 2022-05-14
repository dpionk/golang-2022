// Write a function that checks if a given string (case insensitive) is a palindrome.

package main

import "fmt"

func IsPalindrome(str string) bool {
	if len(str) < 2 {
		return true
	} else if str[0] != str[len(str)-1] {
		return false
	} else {
		return IsPalindrome(str[1:(len(str) - 1)])
	}
}

func main() {
	fmt.Println(IsPalindrome("abba"))
	fmt.Println(IsPalindrome("abbc"))
	fmt.Println(IsPalindrome("kajak"))
}
