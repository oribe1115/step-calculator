package lib

import (
	"fmt"
	"strconv"
)

type Formula struct {
	List []*Token
}

func InitFormula(line string) (*Formula, error) {
	formula := &Formula{}
	formula.List = make([]*Token, 0)
	for len(line) > 0 {
		var token *Token
		var err error

		if isDigit(line[0]) {
			token, line, err = ReadNumber(line)
			if err != nil {
				return nil, err
			}
		} else {
			switch line[0] {
			case '+':
				token, line = ReadPlus(line)
				break
			case '-':
				token, line = ReadMinus(line)
				break
			case '*':
				token, line = ReadMultiply(line)
				break
			case '/':
				token, line = ReadDivision(line)
				break
			case '(':
				token, line, err = ReadBracket(line)
				if err != nil {
					return nil, err
				}
				break
			default:
				return nil, fmt.Errorf("invalid char: %c", line[0])
			}
		}
		formula.List = append(formula.List, token)
	}

	if err := formula.CheckFormat(); err != nil {
		return formula, err
	}

	return formula, nil
}

func (f *Formula) PrintList() {
	for _, token := range f.List {
		fmt.Printf("%v ", token)
	}
	fmt.Print("\n")
}

func (f *Formula) Calc() (float64, error) {
	err := f.calcMulDiv()
	if err != nil {
		return 0, err
	}
	result, err := f.calcPlusMinus()
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (f *Formula) calcPlusMinus() (float64, error) {
	result := float64(0)
	list := make([]*Token, 0)
	list = append(list, CreatePlusToken())
	list = append(list, f.List...)

	for i, token := range list {
		if token.Type == TypePlus || token.Type == TypeMinus {
			num, err := list[i+1].GetNumber()
			if err != nil {
				return 0, err
			}

			if token.Type == TypePlus {
				result += num
			} else {
				result -= num
			}
		}
	}

	return result, nil
}

func (f *Formula) calcMulDiv() error {
	list := make([]*Token, 0)

	for i := 0; i < len(f.List); {
		switch f.List[i].Type {
		case TypeMultiply:
			tmp, err := list[len(list)-1].GetNumber()
			if err != nil {
				return err
			}
			multiplier, err := f.List[i+1].GetNumber()
			if err != nil {
				return err
			}
			tmp *= multiplier
			list[len(list)-1] = CreateNumberToken(tmp)
			i += 2
			break
		case TypeDivision:
			tmp, err := list[len(list)-1].GetNumber()
			if err != nil {
				return err
			}
			divisor, err := f.List[i+1].GetNumber()
			if err != nil {
				return err
			} else if divisor == 0 {
				return fmt.Errorf("found division by zero")
			}
			tmp /= divisor
			list[len(list)-1] = CreateNumberToken(tmp)
			i += 2
			break
		default:
			list = append(list, f.List[i])
			i++
		}
	}

	f.List = list
	return nil
}

// CheckFormat 問題なく計算が行える構成になっているか確認
func (f *Formula) CheckFormat() error {
	if len(f.List) == 0 || len(f.List)%2 != 1 {
		return fmt.Errorf("invalid list length: %d", len(f.List))
	}

	for i := 0; i < len(f.List); i += 2 {
		tokenType := f.List[i].Type
		if !(tokenType == TypeNumber || tokenType == TypeBracket) {
			return fmt.Errorf("index %d should be %s but %s", i, TypeNumber, tokenType)
		}
	}

	for i := 1; i < len(f.List); i += 2 {
		tokenType := f.List[i].Type
		if tokenType == TypeNumber || tokenType == TypeBracket {
			return fmt.Errorf("index %d should be operator but %s", i, tokenType)
		}
	}

	return nil
}

func isDigit(char byte) bool {
	_, err := strconv.Atoi(string(char))
	return err == nil
}
