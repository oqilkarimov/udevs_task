package main

import "fmt"

// need to improve
func fib(number int) int {
	if number == 0 || number == 1 {
		return number
	}
	return fib(number-2) + fib(number-1)
}

func fizzbuzz() {
	for i := 1; i <= 100; i++ {

		if i%3 == 0 {
			fmt.Printf("Fizz")
		}
		if i%5 == 0 {
			fmt.Printf("Buzz")
		}

		if i%3 != 0 && i%5 != 0 {
			fmt.Printf("%d", i)
		}

		fmt.Printf("\n")

	}
}

func IsPalindrome(word string) bool {

	for i := 0; i < len(word)/2; i++ {

		if word[i] != word[len(word)-1-i] {
			return false
		}

	}

	return true
}

func IsOdd(number int) bool {
	if number%2 == 0 {
		return true
	}

	return false
}
