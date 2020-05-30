package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input   string
	Expect  float64
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

		{"10-5", 5, false},
		{"0-5", -5, false},
		{"1-0.25", 0.75, false},
		{"-", 0, true},
		{"1-+3", 0, true},

		{"1+5-0.3", 5.7, false},
		{".3", 0, true},
		{"1..3+", 0, true},
	}

	mulDivCases = []TestCase{
		{"6*2", 12, false},
		{"6/3", 2, false},
		{"1/3", 0.3333333333333333, false},
		{"*", 0, true},
		{"/", 0, true},
		{"3*", 0, true},

		{"1+2*3", 7, false},
		{"4-6/2", 1, false},
		{"6/2*3+1", 10, false},
	}
)

func TestBasicCase(t *testing.T) {
	runTest(t, basicCases)
}

func TestMulDivCase(t *testing.T) {
	runTest(t, mulDivCases)
}

func runTest(t *testing.T, testCases []TestCase) {
	for _, test := range testCases {
		formula, err := InitFormula(test.Input)
		if test.IsError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.Expect, formula.Calc())
		}
	}
}
