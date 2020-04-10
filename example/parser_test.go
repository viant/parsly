package example

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/viant/assertly"
	"testing"
)

func TestNewParser(t *testing.T) {

	useCases := []struct {
		description string
		input       string
		expect      string
		hasError    bool
	}{
		{
			description: "2 operands expression",
			input:       "3 + 8",
			expect: `{
	"LeftOp": {
		"Value": 3
	},
	"Operator": "+",
	"RightOp": {
		"Value": 8
	}
}`,
		},
		//{
		//	description:"3 operands expr ",
		//	input:"3 + 8 * 10",
		//	expect:`{"LeftOp":{"Value":3},"Operator":"+","RightOp":{"Expression":{"LeftOp":{"Value":8},"Operator":"*","RightOp":{"Value":10}}}}`,
		//},
		//{
		//	description:"4 operands expr ",
		//	input:"2 - 3 + 8 * 10",
		//	expect:`{"LeftOp":{"Value":2},"Operator":"-","RightOp":{"Expression":{"LeftOp":{"Value":3},"Operator":"+","RightOp":{"Expression":{"LeftOp":{"Value":8},"Operator":"*","RightOp":{"Value":10}}}}}}`,
		//},
	}

	for _, useCase := range useCases {
		actual, err := Parse([]byte(useCase.input))
		if useCase.hasError {
			assert.NotNil(t, err, useCase.description)
			continue
		}
		if !assert.Nil(t, err, useCase.description) {
			continue
		}
		if !assertly.AssertValues(t, useCase.expect, actual, useCase.description) {
			data, _ := json.Marshal(actual)
			fmt.Printf("%s\n", data)
		}
	}

}
