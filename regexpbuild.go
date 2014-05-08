package regexpbuild

import (
	"fmt"
	"regexp"
	"strings"
)

type Builder interface {
	Build() string
}

type SimpleBuilder struct {
	a string
}

func (b SimpleBuilder) Build() string {
	return b.a
}

func Digit() Builder {
	return SimpleBuilder{`\d`}
}

func NotDigit() Builder {
	return SimpleBuilder{`\D`}
}

func Whitespace() Builder {
	return SimpleBuilder{`\s`}
}

func NotWhitespace() Builder {
	return SimpleBuilder{`\S`}
}

func WordCharacter() Builder {
	return SimpleBuilder{`\w`}
}

func NotWordCharacter() Builder {
	return SimpleBuilder{`\W`}
}

func Literal(s string) Builder {
	re := regexp.MustCompile("[!-/:-@[-`{-~]")
	return SimpleBuilder{re.ReplaceAllStringFunc(s, func(in string) string {
		return `\` + in
	})}
}

func ZeroOrMore(b Builder) Builder {
	return SimpleBuilder{fmt.Sprintf("%s*", b.Build())}
}

func OneOrMore(b Builder) Builder {
	return SimpleBuilder{fmt.Sprintf("%s+", b.Build())}
}

func ZeroOrOne(b Builder) Builder {
	return SimpleBuilder{fmt.Sprintf("%s?", b.Build())}
}

func MinToMax(b Builder, min, max int) Builder {
	return SimpleBuilder{fmt.Sprintf("%s{%d,%d}", b.Build(), min, max)}
}

func ZeroOrMoreLazy(b Builder) Builder {
	return SimpleBuilder{fmt.Sprintf("%s?", ZeroOrMore(b).Build())}
}

func OneOrMoreLazy(b Builder) Builder {
	return SimpleBuilder{fmt.Sprintf("%s?", OneOrMore(b).Build())}
}

func ZeroOrOneLazy(b Builder) Builder {
	return SimpleBuilder{fmt.Sprintf("%s?", ZeroOrOne(b).Build())}
}

func MinToMaxLazy(b Builder, min, max int) Builder {
	return SimpleBuilder{fmt.Sprintf("%s?", MinToMax(b, min, max).Build())}
}

func Group(b Builder) Builder {
	return SimpleBuilder{fmt.Sprintf("(?:%s)", b.Build())}
}

func CapturingGroup(b Builder) Builder {
	return SimpleBuilder{fmt.Sprintf("(%s)", b.Build())}
}

func WordBoundary() Builder {
	return SimpleBuilder{`\b`}
}

func BeginningOfLine() Builder {
	return SimpleBuilder{"^"}
}

func EndOfLine() Builder {
	return SimpleBuilder{"$"}
}

func Or(bb ...Builder) Builder {
	ss := make([]string, len(bb))
	for i := range bb {
		ss[i] = bb[i].Build()
	}

	return SimpleBuilder{strings.Join(ss, "|")}
}

func Sequence(bb ...Builder) Builder {
	ss := make([]string, len(bb))
	for i := range bb {
		ss[i] = bb[i].Build()
	}

	return SimpleBuilder{strings.Join(ss, "")}
}
