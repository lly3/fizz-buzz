package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		name string
		n    int
		want string
	}{
		{
			name: "should return number as string when given the number which cant devide by 3 or 5",
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
			name: "should return 'FizzBuzz' when given the number that can divide by 5 and 3",
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
