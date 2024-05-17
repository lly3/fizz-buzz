package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	// Given
	cases := []struct {
		name string
		n    int
		want string
	}{
		{
			name: "should return '1' when given input equal '1'",
			n:    1,
			want: "1",
		},
		{
			name: "should return '2' when given input equal '2'",
			n:    2,
			want: "2",
		},
		{
			name: "should return 'Fizz' when given input equal '3'",
			n:    3,
			want: "Fizz",
		},
		{
			name: "should return 'Buzz' when given input equal '5'",
			n:    5,
			want: "Buzz",
		},
		{
			name: "should return 'FizzBuzz' when given n equal 15",
			n:    15,
			want: "FizzBuzz",
		},
	}
	// Assert
	for _, c := range cases {
		got := fizzbuzz(c.n)
		if got != c.want {
			t.Errorf("Want '%s', got '%s'", c.want, got)
		}
	}
}
