// Write a function that takes a string of parentheses, and determines if the order of the parentheses is valid.
// The function should return true if the string is valid, and false if it's invalid.

// Examples
// "()"              =>  true
// ")(()))"          =>  false
// "("               =>  false
// "(())((()())())"  =>  true

package main

import "fmt"

func ValidParentheses(parens string) bool {
	openCounter := 0
	closeCounter := 0
	for i := 0; i < len(parens); i++ {
		switch parens[i] {
		case '(':
			openCounter++
		case ')':
			closeCounter++
			if closeCounter > openCounter {
				return false
			}
		}
	}
	if openCounter == closeCounter {
		return true
	}
	return false
}

func main() {
	fmt.Println(ValidParentheses("()"))
	fmt.Println(ValidParentheses(")(()))"))
	fmt.Println(ValidParentheses("("))
	fmt.Println(ValidParentheses("(())((()())())"))
}
