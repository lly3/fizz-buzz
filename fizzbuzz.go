package fizzbuzz

import "fmt"

func fizzbuzz(n int) string {
	if n%3 == 0 {
		return "Fizz"
	}
	return fmt.Sprint(n)
}
