package lib

import (
	"regexp"
	"strconv"
	"strings"
)

type Token struct {
	Type   TokenType
	Number float64
}

type TokenType string

const (
	TypeNumber   TokenType = "NUMBER"
	TypePlus     TokenType = "PLUS"
	TypeMinus    TokenType = "MINUS"
	TypeMultiply TokenType = "MULTIPLY"
	TypeDivision TokenType = "DIVISION"
)

func CreateNumberToken(num float64) *Token {
	return &Token{
		Type:   TypeNumber,
		Number: num,
	}
}

func CreatePlusToken() *Token {
	return &Token{
		Type:   TypePlus,
		Number: 0,
	}
}

func CreateMinusToken() *Token {
	return &Token{
		Type:   TypeMinus,
		Number: 0,
	}
}

func CreateMultiplyToken() *Token {
	return &Token{
		Type:   TypeMultiply,
		Number: 0,
	}
}

func CreateDivisionToken() *Token {
	return &Token{
		Type:   TypeDivision,
		Number: 0,
	}
}

func ReadNumber(line string) (token *Token, remainder string, err error) {
	re := regexp.MustCompile("^\\d+(\\.\\d+)?")
	numString := re.FindString(line)
	num, err := strconv.ParseFloat(numString, 64)
	if err != nil {
		return nil, "", err
	}

	token = CreateNumberToken(num)
	remainder = strings.TrimLeft(line, numString)

	return token, remainder, nil
}

func ReadPlus(line string) (token *Token, remainder string) {
	token = CreatePlusToken()
	remainder = strings.TrimLeft(line, "+")
	return token, remainder
}

func ReadMinus(line string) (token *Token, remainder string) {
	token = CreateMinusToken()
	remainder = strings.TrimLeft(line, "-")
	return token, remainder
}

func ReadMultiply(line string) (token *Token, remainder string) {
	token = CreateMultiplyToken()
	remainder = strings.TrimLeft(line, "*")
	return token, remainder
}

func ReadDivision(line string) (token *Token, remainder string) {
	token = CreateDivisionToken()
	remainder = strings.TrimLeft(line, "/")
	return token, remainder
}

func (t *Token) GetNumber() float64 {
	return t.Number
}
