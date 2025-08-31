package main

import "fmt"

func main() {
	fmt.Println("Factorial of 10 ->", findFactorial(10))
	fmt.Println("Digit Sum of a number 124581902 ->", findDigitSumOfNumber(124581902))
}

func findFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * findFactorial(n-1)
}

func findDigitSumOfNumber(num int) int {
	if num < 10 {
		return num
	}

	return num%10 + findDigitSumOfNumber(num/10)
}
