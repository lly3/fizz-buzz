package fizzbuzz

import "fmt"

func ternary(condition bool, a, b interface{}) interface{} {
	m := map[bool]interface{}{
		true:  a,
		false: b,
	}
	return m[condition]
}

func fizzbuzz(n int) string {
	s := fmt.Sprint(ternary(n%3 == 0, "Fizz", "")) + fmt.Sprint(ternary(n%5 == 0, "Buzz", ""))
	return fmt.Sprint(ternary(s == "", fmt.Sprint(n), s))
}
