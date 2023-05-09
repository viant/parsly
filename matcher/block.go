package matcher

import "github.com/viant/parsly"

type Block struct {
	escape     byte
	begin      byte
	end        byte
	matchQuote bool
}

//TokenMatch matches quoted characters
func (m *Block) Match(cursor *parsly.Cursor) int {
	if m.matchQuote {
		return m.matchQoutes(cursor)
	}
	var matched = 0
	input := cursor.Input
	inputSize := len(input)
	pos := cursor.Pos
	value := input[pos]
	if value != m.begin {
		return 0
	}
	depth := 1
	matched++

	var inQuote byte
	for i := pos + matched; i < inputSize; i++ {
		value = input[i]
		isInQuote := inQuote != 0
		matched++
		switch value {
		case '"', '`', '\'': //quotes
			if matched > 1 {
				if input[i-1] == m.escape { //if escaped quote
					continue
				}
			}
			if !isInQuote {
				inQuote = value
			} else if inQuote == value {
				inQuote = 0
			}

		case m.begin:
			if m.begin == m.end {
				return matched
			}

			if isInQuote {
				continue
			}
			depth++

		case m.end:
			if isInQuote {
				continue
			}
			depth--
			if depth == 0 {
				return matched
			}
		}
	}
	return 0
}

//TokenMatch matches quoted characters
func (m *Block) matchQoutes(cursor *parsly.Cursor) int {
	var matched = 0
	input := cursor.Input
	inputSize := len(input)
	pos := cursor.Pos
	value := input[pos]
	if value != m.begin {
		return 0
	}
	depth := 1
	matched++

	var inQuote byte
	for i := pos + matched; i < inputSize; i++ {
		value = input[i]
		isInQuote := inQuote != 0
		matched++
		switch value {
		case m.begin:
			if i > 0 && input[i-1] == m.escape {
				continue
			}
			if m.begin == m.end {
				return matched
			}

			if isInQuote {
				continue
			}
			depth++

		case m.end:
			if i > 0 && input[i-1] == m.escape {
				continue
			}
			depth--
			if depth == 0 {
				return matched
			}
		}
	}
	return 0
}

func NewBlock(begin, end, escape byte) *Block {
	return &Block{
		escape:     escape,
		begin:      begin,
		end:        end,
		matchQuote: begin == end && (begin == '\'' || begin == '"'),
	}
}
