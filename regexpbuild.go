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

func ZeroOrMore(a Builder) Builder {
	return SimpleBuilder{fmt.Sprintf("%s*", a.Build())}
}

func OneOrMore(a Builder) Builder {
	return SimpleBuilder{fmt.Sprintf("%s+", a.Build())}
}

func ZeroOrOne(a Builder) Builder {
	return SimpleBuilder{fmt.Sprintf("%s?", a.Build())}
}

func MinToMax(a Builder, min, max int) Builder {
	return SimpleBuilder{fmt.Sprintf("%s{%d,%d}", a.Build(), min, max)}
}

func ZeroOrMoreLazy(a Builder) Builder {
	return SimpleBuilder{fmt.Sprintf("%s?", ZeroOrMore(a).Build())}
}

func OneOrMoreLazy(a Builder) Builder {
	return SimpleBuilder{fmt.Sprintf("%s?", OneOrMore(a).Build())}
}

func ZeroOrOneLazy(a Builder) Builder {
	return SimpleBuilder{fmt.Sprintf("%s?", ZeroOrOne(a).Build())}
}

func MinToMaxLazy(a Builder, min, max int) Builder {
	return SimpleBuilder{fmt.Sprintf("%s?", MinToMax(a, min, max).Build())}
}

func Group(s Builder) Builder {
	return SimpleBuilder{fmt.Sprintf("(?:%s)", s.Build())}
}

func CapturingGroup(s Builder) Builder {
	return SimpleBuilder{fmt.Sprintf("(%s)", s.Build())}
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
