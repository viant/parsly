package example


type Expression struct {
	LeftOp   *Operand `json:",omitempty"`
	Operator string
	RightOp  *Operand `json:",omitempty"`
}
