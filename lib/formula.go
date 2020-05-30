package lib

import (
	"fmt"
	"strconv"
)

type Formula struct {
	List []*Token
}

func InitFormula(line string) *Formula {
	formula := &Formula{}
	formula.List = make([]*Token, 0)
	for len(line) > 0 {
		var token *Token
		var err error

		if isDigit(line[0]) {
			token, line, err = ReadNumber(line)
			if err != nil {
				panic(err)
			}
		} else {
			switch line[0] {
			case '+':
				token, line = ReadPlus(line)
				break
			default:
				fmt.Printf("invalid char: %c", line[0])
				return nil
			}
		}
		formula.List = append(formula.List, token)
	}

	return formula
}

func (f *Formula) PrintList() {
	for _, token := range f.List {
		fmt.Printf("%v ", token)
	}
	fmt.Print("\n")
}

func isDigit(char byte) bool {
	_, err := strconv.Atoi(string(char))
	return err == nil
}
