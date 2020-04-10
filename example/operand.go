package example

type Operand struct {
	Value      *float64    `json:",omitempty"`
	Expression *Expression `json:",omitempty"`
}

//NewValue creates value  operand
func NewValue(value float64) *Operand {
	return &Operand{Value: &value}
}
