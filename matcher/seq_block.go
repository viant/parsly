package matcher

import "github.com/viant/parsly"

type SeqBlock struct {
	begin []byte
	end   []byte
}

//TokenMatch matches quoted characters
func (m *SeqBlock) Match(cursor *parsly.Cursor) int {
	var matched = 0
	input := cursor.Input
	inputSize := len(input)
	pos := cursor.Pos

	beginSize := len(m.begin)
	endSize := len(m.end)
	value := input[pos]
	if !MatchFold(m.begin, input, 0, pos) {
		return 0
	}

	depth := 1
	matched += beginSize

	var inQuote byte
	for i := pos + matched; i < inputSize; i++ {
		value = input[i]
		isInQuote := inQuote != 0
		matched++
		switch value {
		case '"', '`', '\'': //quotes
			if !isInQuote {
				inQuote = value
			} else if inQuote == value {
				inQuote = 0
			}

		case m.end[0]:
			if !MatchFold(m.end, input, 0, i) {
				continue
			}
			i += endSize - 1
			matched += endSize - 1
			if isInQuote {
				continue
			}
			depth--
			if depth == 0 {
				return matched
			}

			if m.end[0] != m.begin[0] {
				continue
			}
			fallthrough
		case m.begin[0]:
			if !MatchFold(m.begin, input, 0, i) {
				continue
			}
			i += beginSize - 1
			matched += beginSize - 1
			if isInQuote {
				continue
			}
			depth++

		}

	}
	return 0
}

func NewSeqBlock(begin, end string) *SeqBlock {
	return &SeqBlock{
		begin: []byte(begin),
		end:   []byte(end),
	}
}
