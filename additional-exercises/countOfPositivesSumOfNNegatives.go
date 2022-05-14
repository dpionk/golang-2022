// Given an array of integers.
// Return an array, where the first element is the count of positives numbers and the second element is sum of negative numbers. 0 is neither positive nor negative.
// If the input is an empty array or is null, return an empty array.
// Example
// For input [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15], you should return [10, -65].

package main

import "fmt"

func CountPositivesSumNegatives(numbers []int) [2]int {
	var res [2]int
	positiveNumbersCounter := 0
	negativeNumbersSum := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 0 {
			positiveNumbersCounter++
		} else if numbers[i] < 0 {
			negativeNumbersSum += numbers[i]
		}
	}
	res[0], res[1] = positiveNumbersCounter, negativeNumbersSum
	return res
}

func main() {
	fmt.Println(CountPositivesSumNegatives([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}))
}
