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
		{"1", 1, false},
		{"10", 10, false},
		{"1.02", 1.02, false},
		{"0.004", 0.004, false},
		{"+", 0, true},
		{"1+1", 2, false},
		{"1+", 0, true},
		{"1++", 0, true},
		{"12+1.03", 13.03, false},
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
