// Write a function that accepts an array of 10 integers (between 0 and 9),
// that returns a string of those numbers in the form of a phone number.

// Example
// CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})  // returns "(123) 456-7890"
// The returned format must be correct in order to complete this challenge.
// Don't forget the space after the closing parentheses!

package main

import (
	"fmt"
	"strconv"
)

func CreatePhoneNumber(numbers [10]uint) string {
	phoneNumber := "("
	for i := 0; i < len(numbers); i++ {
		switch {
		case i < 2 || (i > 2 && i < 5) || (i > 5 && i < len(numbers)):
			phoneNumber += strconv.Itoa(int(numbers[i]))
		case i == 2:
			phoneNumber += strconv.Itoa(int(numbers[i])) + ") "
		case i == 5:
			phoneNumber += strconv.Itoa(int(numbers[i])) + "-"
		}
	}
	return phoneNumber
}

func main() {
	fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}
