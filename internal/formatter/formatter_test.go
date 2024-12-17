package formatter

import "testing"

func TestCompositeFormatter(t *testing.T) {
    f := NewCompositeFormatter(NewChain(NewDefaultStrategy()))

    cases := []struct {
        in   int
        want string
    }{
        {1, "1"},
        {3, "Fizz"},
        {5, "Buzz"},
        {15, "FizzBuzz"},
    }

    for _, c := range cases {
        got := f.Format(c.in)
        if got != c.want {
            t.Errorf("Format(%d) = %s; want %s", c.in, got, c.want)
        }
    }
}
