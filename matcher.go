package golexer

import (
	"bytes"
	"errors"

	"unicode"
)

type TokenMatcher interface {
	Match(*Tokenizer) (*Token, error)
	ID() int
}

type baseMatcher struct {
	id int
}

func (self *baseMatcher) ID() int {
	return self.id
}

// 未知字符
type UnknownMatcher struct {
	baseMatcher
}

func (self *UnknownMatcher) Match(tz *Tokenizer) (*Token, error) {

	if tz.Current() == 0 {
		return nil, nil
	}

	begin := tz.Index()

	tz.ConsumeOne()

	return NewToken(self, tz, tz.StringRange(begin, tz.Index())), nil
}

func NewUnknownMatcher(id int) TokenMatcher {
	return &UnknownMatcher{
		baseMatcher{id},
	}
}

// 空白字符
type WhiteSpaceMatcher struct {
	baseMatcher
}

func isWhiteSpace(c rune) bool {
	return c == ' ' || c == '\t'
}

func (self *WhiteSpaceMatcher) Match(tz *Tokenizer) (*Token, error) {

	var count int
	for {

		c := tz.Peek(count)

		if !isWhiteSpace(c) {
			break
		}

		count++

	}

	if count == 0 {
		return nil, nil
	}

	tz.ConsumeMulti(count)

	return NewToken(self, tz, ""), nil
}

func NewWhiteSpaceMatcher(id int) TokenMatcher {
	return &WhiteSpaceMatcher{
		baseMatcher{id},
	}
}

// 行结束
type LineEndMatcher struct {
	baseMatcher
}

func (self *LineEndMatcher) Match(tz *Tokenizer) (*Token, error) {

	var count int
	for {

		c := tz.Peek(count)

		if c == '\n' {
			tz.increaseLine()
		} else if c == '\r' {

		} else {
			break
		}

		count++

	}

	if count == 0 {
		return nil, nil
	}

	tz.ConsumeMulti(count)

	return NewToken(self, tz, ""), nil
}

func NewLineEndMatcher(id int) TokenMatcher {
	return &LineEndMatcher{
		baseMatcher{id},
	}
}

// 整形，浮点数
type NumeralMatcher struct {
	baseMatcher
}

func (self *NumeralMatcher) Match(tz *Tokenizer) (*Token, error) {

	if !unicode.IsDigit(tz.Current()) && tz.Current() != '-' {
		return nil, nil
	}

	begin := tz.Index()

	var maybeFloat bool

	for {

		tz.ConsumeOne()

		if !unicode.IsDigit(tz.Current()) {

			if tz.Current() == '.' {
				maybeFloat = true
			}

			break
		}

	}

	if maybeFloat {
		for i := 0; ; i++ {

			tz.ConsumeOne()

			if !unicode.IsDigit(tz.Current()) {

				// .之后的第一个字符居然不是数字
				if i == 0 {
					return nil, errors.New("invalid numeral")
				}

				break

			}

		}
	}

	return NewToken(self, tz, tz.StringRange(begin, tz.Index())), nil
}

func NewNumeralMatcher(id int) TokenMatcher {
	return &NumeralMatcher{
		baseMatcher{id},
	}
}

// 字符串
type StringMatcher struct {
	baseMatcher
	builder bytes.Buffer
}

func (self *StringMatcher) Match(tz *Tokenizer) (*Token, error) {

	if tz.Current() != '"' && tz.Current() != '\'' {
		return nil, nil
	}

	tz.ConsumeOne()

	var escaping bool

	self.builder.Reset()

	for {

		if escaping {
			switch tz.Current() {
			case 'n':
				self.builder.WriteRune('\n')
			case 'r':
				self.builder.WriteRune('\r')
			default:
				self.builder.WriteRune('\\')
				self.builder.WriteRune(tz.Current())
			}

			escaping = false
		} else {
			if tz.Current() == '\\' {
				escaping = true
			} else {
				self.builder.WriteRune(tz.Current())
			}
		}

		tz.ConsumeOne()

		if tz.Current() == '\n' ||
			tz.Current() == 0 ||
			tz.Current() == '"' ||
			tz.Current() == '\'' {
			break
		}

	}

	tz.ConsumeOne()

	return NewToken(self, tz, self.builder.String()), nil
}

func NewStringMatcher(id int) TokenMatcher {
	return &StringMatcher{
		baseMatcher: baseMatcher{id},
	}
}

// 标识符
type IdentifierMatcher struct {
	baseMatcher
}

func (self *IdentifierMatcher) Match(tz *Tokenizer) (*Token, error) {

	if !unicode.IsLetter(tz.Current()) && tz.Current() != '_' {
		return nil, nil
	}

	begin := tz.Index()

	for {

		tz.ConsumeOne()

		if !unicode.IsLetter(tz.Current()) && tz.Current() != '_' {
			break
		}

	}

	return NewToken(self, tz, tz.StringRange(begin, tz.index)), nil
}

func NewIdentifierMatcher(id int) TokenMatcher {
	return &IdentifierMatcher{
		baseMatcher{id},
	}
}

// 操作符，分隔符，关键字
type SignMatcher struct {
	baseMatcher
	word []rune
}

func (self *SignMatcher) Match(tz *Tokenizer) (*Token, error) {

	if (tz.Count() - tz.Index()) < len(self.word) {
		return nil, nil
	}

	for i, c := range self.word {

		if tz.Peek(i) != c {
			return nil, nil
		}

	}

	tz.ConsumeMulti(len(self.word))

	return NewToken(self, tz, string(self.word)), nil
}

func NewSignMatcher(id int, word string) TokenMatcher {
	return &SignMatcher{
		baseMatcher: baseMatcher{id},
		word:        []rune(word),
	}
}

// #开头的行注释
type UnixStyleCommentMatcher struct {
	baseMatcher
}

func (self *UnixStyleCommentMatcher) Match(tz *Tokenizer) (*Token, error) {
	if tz.Current() != '#' {
		return nil, nil
	}

	tz.ConsumeOne()

	begin := tz.Index()

	for {

		tz.ConsumeOne()

		if tz.Current() == '\n' || tz.Current() == 0 {
			break
		}

	}

	return NewToken(self, tz, tz.StringRange(begin, tz.index)), nil
}

func NewUnixStyleCommentMatcher(id int) TokenMatcher {
	return &UnixStyleCommentMatcher{
		baseMatcher{id},
	}
}
