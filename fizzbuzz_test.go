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
			name:      "should return a when given contiditon true",
			condition: true,
			a:         "dummy a",
			b:         "dummy b",
			want:      "dummy a",
		},
		{
			name:      "should return b when given contiditon false",
			condition: false,
			a:         "dummy a",
			b:         "dummy b",
			want:      "dummy b",
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
			name: "should return 1 when given the number that cant devide by 3 or 5",
			n:    1,
			want: "1",
		},
		{
			name: "should return 'Fizz' when given the number that can devide by 3",
			n:    3,
			want: "Fizz",
		},
	}

	for _, c := range cases {
		got := fizzbuzz(c.n)
		if got != c.want {
			t.Errorf("want: %v, got: %v", c.want, got)
		}
	}
}
