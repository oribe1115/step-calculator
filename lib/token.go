package lib

import (
	"fmt"
	"regexp"
	"strconv"
)

type Token struct {
	Type      TokenType
	Number    float64
	InBracket *Formula
}

type TokenType string

const (
	TypeNumber   TokenType = "NUMBER"
	TypePlus     TokenType = "PLUS"
	TypeMinus    TokenType = "MINUS"
	TypeMultiply TokenType = "MULTIPLY"
	TypeDivision TokenType = "DIVISION"
	TypeBracket  TokenType = "BRACKET"
)

func CreateNumberToken(num float64) *Token {
	return &Token{
		Type:      TypeNumber,
		Number:    num,
		InBracket: nil,
	}
}

func CreatePlusToken() *Token {
	return &Token{
		Type:      TypePlus,
		Number:    0,
		InBracket: nil,
	}
}

func CreateMinusToken() *Token {
	return &Token{
		Type:      TypeMinus,
		Number:    0,
		InBracket: nil,
	}
}

func CreateMultiplyToken() *Token {
	return &Token{
		Type:      TypeMultiply,
		Number:    0,
		InBracket: nil,
	}
}

func CreateDivisionToken() *Token {
	return &Token{
		Type:      TypeDivision,
		Number:    0,
		InBracket: nil,
	}
}

func CreateBracketToken(inBracketStr string) (*Token, error) {
	inBracket, err := InitFormula(inBracketStr)
	if err != nil {
		return nil, err
	}

	token := &Token{
		Type:      TypeBracket,
		Number:    0,
		InBracket: inBracket,
	}

	return token, nil
}

func ReadNumber(line string) (token *Token, remainder string, err error) {
	re := regexp.MustCompile("^\\d+(\\.\\d+)?")
	numString := re.FindString(line)
	num, err := strconv.ParseFloat(numString, 64)
	if err != nil {
		return nil, "", err
	}

	token = CreateNumberToken(num)
	remainder = line[len(numString):]

	return token, remainder, nil
}

func ReadPlus(line string) (token *Token, remainder string) {
	token = CreatePlusToken()
	remainder = line[1:]
	return token, remainder
}

func ReadMinus(line string) (token *Token, remainder string) {
	token = CreateMinusToken()
	remainder = line[1:]
	return token, remainder
}

func ReadMultiply(line string) (token *Token, remainder string) {
	token = CreateMultiplyToken()
	remainder = line[1:]
	return token, remainder
}

func ReadDivision(line string) (token *Token, remainder string) {
	token = CreateDivisionToken()
	remainder = line[1:]
	return token, remainder
}

func ReadBracket(line string) (token *Token, remainder string, err error) {
	leftCount := 1
	rightCount := 0
	rightBracketIndex := 0

	for i := 1; i < len(line); i++ {
		if line[i] == '(' {
			leftCount++
		} else if line[i] == ')' {
			rightCount++
		}

		if leftCount == rightCount {
			rightBracketIndex = i
			break
		}
	}

	if rightBracketIndex == 0 {
		return nil, line, fmt.Errorf("faild to find right bracket: %s", line)
	}

	token, err = CreateBracketToken(line[1:rightBracketIndex])
	if err != nil {
		return nil, line, err
	}

	remainder = line[rightBracketIndex+1:]

	return token, remainder, nil
}

func (t *Token) GetNumber() float64 {
	if t.Type == TypeBracket {
		return t.InBracket.Calc()
	}
	return t.Number
}
