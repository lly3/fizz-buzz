package fizzbuzz

import "fmt"

type FizzBuzz struct{}

func NewFizzBuzz() FizzBuzz {
	return FizzBuzz{}
}

func (fb FizzBuzz) Run(n int) string {
	return fmt.Sprint(n)
}

func fizzbuzz(n int) string {
	fb := NewFizzBuzz()
	return fb.Run(n)
}
