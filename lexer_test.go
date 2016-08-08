package golexer

import (
	"testing"
)

// 自定义的token id
const (
	Token_Unknown = iota
	Token_Numeral
	Token_String
	Token_WhiteSpace
	Token_LineEnd
	Token_UnixStyleComment
	Token_Identifier
	Token_Go
	Token_Semicolon
)

func TestLexer(t *testing.T) {

	l := NewLexer()

	// 匹配顺序从高到低

	l.AddMatcher(NewNumeralMatcher(Token_Numeral))
	l.AddMatcher(NewStringMatcher(Token_String))

	l.AddIgnoreMatcher(NewWhiteSpaceMatcher(Token_WhiteSpace))
	l.AddIgnoreMatcher(NewLineEndMatcher(Token_LineEnd))
	l.AddIgnoreMatcher(NewUnixStyleCommentMatcher(Token_UnixStyleComment))

	l.AddMatcher(NewSignMatcher(Token_Semicolon, ";"))
	l.AddMatcher(NewSignMatcher(Token_Go, "go"))

	l.AddMatcher(NewIdentifierMatcher(Token_Identifier))

	l.AddMatcher(NewUnknownMatcher(Token_Unknown))

	l.Start(`"a" 
	123.3;
	go
	_id # comment
	;
	'b'
	
	
	`)

	for {

		tk, err := l.Read()

		if err != nil {
			t.Error(err)
			break
		}

		if tk == nil {
			break
		}

		t.Log(tk.String())
	}

}
