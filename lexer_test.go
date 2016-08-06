package golexer

import (
	"testing"
)

func TestLexer(t *testing.T) {

	l := NewLexer()

	l.AddMatcher(new(UnknownMatcher))

	l.Start("哈哈 x b")

	for {

		tk := l.Read()
		if tk == nil {
			break
		}

		t.Log(tk.String())
	}

}
