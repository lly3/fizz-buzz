package fizzbuzz

import "fmt"

type Fizz struct{}

func NewFizz() Fizz {
	return Fizz{}
}

func (f Fizz) Run(n int) string {
	if n%3 == 0 {
		return "Fizz"
	}
	return ""
}

type FizzBuzz struct {
	f Fizz
}

func NewFizzBuzz(f Fizz) FizzBuzz {
	return FizzBuzz{f}
}

func (fb FizzBuzz) Run(n int) string {
	s := fb.f.Run(n)
	if s == "" {
		return fmt.Sprint(n)
	}
	return s
}

func fizzbuzz(n int) string {
	f := NewFizz()
	fb := NewFizzBuzz(f)
	return fb.Run(n)
}
