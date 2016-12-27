package golexer

import (
	"testing"
)

// 自定义的token id
const (
	Token_EOF = iota
	Token_Unknown
	Token_Numeral
	Token_String
	Token_WhiteSpace
	Token_LineEnd
	Token_UnixStyleComment
	Token_Identifier
	Token_Dot
	Token_Go
	Token_XX
	Token_Every
	Token_Week
	Token_Semicolon
)

type CustomParser struct {
	*Parser
}

func NewCustomParser() *CustomParser {

	l := NewLexer()

	// 匹配顺序从高到低

	l.AddMatcher(NewNumeralMatcher(Token_Numeral))
	l.AddMatcher(NewStringMatcher(Token_String))

	l.AddIgnoreMatcher(NewWhiteSpaceMatcher(Token_WhiteSpace))
	l.AddIgnoreMatcher(NewLineEndMatcher(Token_LineEnd))
	l.AddIgnoreMatcher(NewUnixStyleCommentMatcher(Token_UnixStyleComment))

	l.AddMatcher(NewSignMatcher(Token_Semicolon, ";"))
	l.AddMatcher(NewSignMatcher(Token_Dot, "."))
	l.AddMatcher(NewKeywordMatcher(Token_Go, "go"))
	l.AddMatcher(NewKeywordMatcher(Token_XX, "xx"))
	l.AddMatcher(NewKeywordMatcher(Token_Every, "每"))
	l.AddMatcher(NewKeywordMatcher(Token_Week, "周"))

	l.AddMatcher(NewIdentifierMatcher(Token_Identifier))

	l.AddMatcher(NewUnknownMatcher(Token_Unknown))

	return &CustomParser{
		Parser: NewParser(l),
	}
}

func TestParser(t *testing.T) {

	p := NewCustomParser()

	defer ErrorCatcher(func(err error) {

		t.Error(err.Error())

	})

	p.Lexer().Start(`"a"
		123.3;
		-1
		gonew.xx
		_id # comment
		每周
		"\'\""
		;
		'b'

		`)

	p.NextToken()

	for p.TokenID() != 0 {

		t.Logf("id: %d value:'%s' name: %s", p.TokenID(), p.TokenValue(), p.MatcherName())

		p.NextToken()

	}

}
