package golexer

import (
	"testing"
)

func TestLexer(t *testing.T) {

	l := NewLexer()

	// 匹配顺序从高到低

	l.AddMatcher(new(NumeralMatcher))
	l.AddMatcher(new(StringMatcher))

	l.AddIgnoreMatcher(new(WhitespaceMatcher))
	l.AddIgnoreMatcher(new(LineEndMatcher))
	l.AddIgnoreMatcher(new(UnixStyleCommentMatcher))

	l.AddMatcher(NewSignMatcher(";"))

	l.AddMatcher(NewSignMatcher("go"))

	l.AddMatcher(new(IdentifierMatcher))

	l.AddMatcher(new(UnknownMatcher))

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
