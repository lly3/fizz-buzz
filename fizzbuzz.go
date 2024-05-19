package fizzbuzz

import "fmt"

func ternary(condition bool, a, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

func fizzbuzz(n int) string {
	s := fmt.Sprint(ternary(n%3 == 0, "Fizz", "")) + fmt.Sprint(ternary(n%5 == 0, "Buzz", ""))
	return fmt.Sprint(ternary(s == "", fmt.Sprint(n), s))
}
