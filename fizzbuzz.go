package fizzbuzz

import "fmt"

func ternary(condition bool, a, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

func fizzbuzz(n int) string {
	return fmt.Sprint(n)
}
