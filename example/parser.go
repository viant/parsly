package example


import (
	"github.com/viant/parsly"
)


func Parse(input []byte) (root *Expression, err error) {
	cursor := parsly.NewCursor("", input, 0)
	root = &Expression{}
	expression := root
	matched := cursor.MatchAfterOptional(Whitespace, Number)
	if matched.Code != Number.Code {
		return nil, cursor.NewError(Number)
	}
	value, _ := matched.Float(cursor)
	expression.LeftOp = NewValue(value)

	for ; ; {
		matched = cursor.MatchAfterOptional(Whitespace, Factor, Term)
		if matched.Code == parsly.EOF {
			break
		}
		operator := matched.Text(cursor)
		if expression.Operator != "" {
			expression.RightOp = &Operand{Expression: &Expression{LeftOp: expression.RightOp}}
			expression = expression.RightOp.Expression
		}
		expression.Operator = operator

		matched := cursor.MatchAfterOptional(Whitespace, Number)
		if matched.Code != Number.Code {
			return nil, cursor.NewError(Number)
		}
		value, _ := matched.Float(cursor)
		expression.RightOp = NewValue(value)
	}
	return root, nil
}
