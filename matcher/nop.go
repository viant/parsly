package matcher

import "github.com/viant/parsly"

type Nop struct {
}

func (c *Nop) Match(cursor *parsly.Cursor) (matched int) {
	return 0
}

//NewNop never matcher
func NewNop() *Nop {
	return &Nop{}
}
