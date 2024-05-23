package fizzbuzz

import "testing"

func TestFizz(t *testing.T) {
	cases := []struct {
		name string
		n    int
		want string
	}{
		{
			name: "should return 'Fizz' when given the number that diviable by 3",
			n:    3,
			want: "Fizz",
		},
		{
			name: "should return 'Fizz' when given the number that diviable by 3",
			n:    2,
			want: "",
		},
	}

	f := NewFizz()

	for _, c := range cases {
		got := f.Run(c.n)
		if got != c.want {
			t.Errorf("want: %v, got: %v", c.want, got)
		}
	}
}

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		name string
		n    int
		want string
	}{
		{
			name: "should return number as string when given the number that can't divide by 3 or 5",
			n:    1,
			want: "1",
		},
		{
			name: "should return 'Fizz' when given the number that can divide by 3",
			n:    3,
			want: "Fizz",
		},
		{
			name: "should return 'Buzz' when given the number that can divide by 5",
			n:    5,
			want: "Buzz",
		},
	}

	for _, c := range cases {
		got := fizzbuzz(c.n)
		if got != c.want {
			t.Errorf("want: %v, got: %v", c.want, got)
		}
	}
}
