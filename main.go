package main

import (
	"os"

	"github.com/01-edu/z01"
)

// isPrime checks if a number is prime.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// sumPrimes calculates the sum of all prime numbers less than or equal to n.
func sumPrimes(n int) int {
	sum := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	return sum
}

// atoi converts a string to an integer without using strconv.
func atoi(s string) (int, bool) {
	result := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0, false
		}
		result = result*10 + int(r-'0')
	}
	return result, true
}

// itoa converts an integer to a string without using strconv.
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		result = string(rune('0'+(n%10))) + result
		n /= 10
	}
	return result
}

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	num, ok := atoi(os.Args[1])
	if !ok || num < 1 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	sum := sumPrimes(num)
	for _, r := range itoa(sum) {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
