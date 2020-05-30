package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input   string
	Result  float64
	IsError bool
}

var (
	basicCases = []TestCase{
		{Input: "1+1", Result: 2, IsError: false},
	}
)

func TestBasicCase(t *testing.T) {
	for _, test := range basicCases {
		formula, err := InitFormula(test.Input)
		if test.IsError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.Result, formula.Calc())
		}
	}
}
