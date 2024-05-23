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

type Buzz struct{}

func NewBuzz() Buzz {
	return Buzz{}
}

func (b Buzz) Run(n int) string {
	if n%5 == 0 {
		return "Buzz"
	}
	return ""
}

type FizzBuzz struct {
	f Fizz
	b Buzz
}

func NewFizzBuzz(f Fizz, b Buzz) FizzBuzz {
	return FizzBuzz{f, b}
}

func (fb FizzBuzz) Run(n int) string {
	s := fb.f.Run(n) + fb.b.Run(n)
	if s == "" {
		return fmt.Sprint(n)
	}
	return s
}

func fizzbuzz(n int) string {
	f := NewFizz()
	b := NewBuzz()
	fb := NewFizzBuzz(f, b)
	return fb.Run(n)
}
