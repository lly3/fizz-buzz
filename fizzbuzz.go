package fizzbuzz

import "fmt"

func fizzbuzz(n int) string {
	if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprint(n)
}
