package regexpbuild

import (
	"fmt"
	"regexp"
	"strings"
)

func Digit() string {
	return "\\d"
}

func NotDigit() string {
	return "\\D"
}

func Whitespace() string {
	return "\\s"
}

func NotWhitespace() string {
	return "\\S"
}

func WordCharacter() string {
	return "\\w"
}

func NotWordCharacter() string {
	return "\\W"
}

func Literal(s string) string {
	re := regexp.MustCompile("[!-/:-@[-`{-~]")
	return re.ReplaceAllStringFunc(s, func(in string) string {
		return "\\" + in
	})
}

func ZeroOrMore(a string) string {
	return fmt.Sprintf("%s*", a)
}

func OneOrMore(a string) string {
	return fmt.Sprintf("%s+", a)
}

func ZeroOrOne(a string) string {
	return fmt.Sprintf("%s?", a)
}

func MinToMax(a string, min, max int) string {
	return fmt.Sprintf("%s{%d,%d}", a, min, max)
}

func ZeroOrMoreLazy(a string) string {
	return fmt.Sprintf("%s?", ZeroOrMore(a))
}

func OneOrMoreLazy(a string) string {
	return fmt.Sprintf("%s?", OneOrMore(a))
}

func ZeroOrOneLazy(a string) string {
	return fmt.Sprintf("%s?", ZeroOrOne(a))
}

func MinToMaxLazy(a string, min, max int) string {
	return fmt.Sprintf("%s?", MinToMax(a, min, max))
}

func Group(s string) string {
	return fmt.Sprintf("(?:%s)", s)
}

func CapturingGroup(s string) string {
	return fmt.Sprintf("(%s)", s)
}

func WordBoundary() string {
	return "\\b"
}

func BeginningOfLine() string {
	return "^"
}

func EndOfLine() string {
	return "$"
}

func Or(ss ...string) string {
	return strings.Join(ss, "|")
}
