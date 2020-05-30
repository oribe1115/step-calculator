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
	TypeNumber TokenType = "NUMBER"
	TypePlus   TokenType = "PLUS"
)

func CreateNumberToken(num float64) *Token {
	return &Token{
		Type:   TypeNumber,
		Number: num,
	}
}

func ReadNumber(line string) (num float64, remainder string, err error) {
	re := regexp.MustCompile("^\\d+\\.\\d")
	numString := re.FindString(line)
	num, err = strconv.ParseFloat(numString, 64)
	if err != nil {
		return 0, "", err
	}

	remainder = strings.TrimLeft(line, numString)

	return num, remainder, nil
}
