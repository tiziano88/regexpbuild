package regexpbuild

import (
	"testing"
)

type testPair struct {
	regexp   Builder
	expected string
}

var tests = []testPair{
	{
		MinToMax(Whitespace(), 3, 10),
		`\s{3,10}`,
	},
	{
		Sequence(
			OneOrMore(WordCharacter()),
			Literal("@"),
			OneOrMore(WordCharacter()),
			Literal("."),
			OneOrMore(WordCharacter())),
		`\w+\@\w+\.\w+`,
	},
	{
		Literal("www.google.com"),
		`www\.google\.com`,
	},
	{
		Sequence(
			MinToMax(Digit(), 1, 3), Literal("."),
			MinToMax(Digit(), 1, 3), Literal("."),
			MinToMax(Digit(), 1, 3), Literal("."),
			MinToMax(Digit(), 1, 3)),
		`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`,
	},
}

func TestRegexp(t *testing.T) {
	for i, pair := range tests {
		if pair.regexp.Build() != pair.expected {
			t.Errorf("For pair #%d expected %s got %s", i, pair.expected, pair.regexp.Build())
		}
	}
}
