package regexpbuild

import (
	"testing"
)

type testPair struct {
	regexp   string
	expected string
}

var tests = []testPair{
	{MinToMax(Digit(), 3, 10), `\d{3,10}`},
	{Literal("www.google.com"), `www\.google\.com`},
}

func TestRegexp(t *testing.T) {
	for i, pair := range tests {
		if pair.regexp != pair.expected {
			t.Errorf("For pair #%d expected %s got %s", i, pair.expected, pair.regexp)
		}
	}
}
