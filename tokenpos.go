package golexer

import (
	"fmt"
)

type TokenPos struct {
	Line       int
	Col        int
	SourceName string
}

func (self TokenPos) String() string {
	return fmt.Sprintf("%s(%d:%d)", self.SourceName, self.Line, self.Col)
}
