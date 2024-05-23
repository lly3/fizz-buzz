package fizzbuzz

import "testing"

func TestTernary(t *testing.T) {
	cases := []struct {
		name      string
		condition bool
		a         interface{}
		b         interface{}
		want      interface{}
	}{
		{
			name:      "should return a when given true condition",
			condition: true,
			a:         "a",
			b:         "b",
			want:      "a",
		},
		{
			name:      "should return b when given false condition",
			condition: false,
			a:         "a",
			b:         "b",
			want:      "b",
		},
	}

	for _, c := range cases {
		got := ternary(c.condition, c.a, c.b)
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
		{
			name: "should return 'Buzz' when given the number that can divide by 5",
			n:    15,
			want: "FizzBuzz",
		},
	}

	for _, c := range cases {
		got := fizzbuzz(c.n)
		if got != c.want {
			t.Errorf("want: %v, got: %v", c.want, got)
		}
	}
}
