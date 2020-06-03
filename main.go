package main

import (
	"fmt"

	"github.com/oribe1115/step-calculator/lib"
)

func main() {

	lib.InitStdin()

	fmt.Printf("formula: ")
	input := lib.ReadLine()

	formula, err := lib.InitFormula(input)
	if err != nil {
		fmt.Println(err)
		if formula != nil {
			formula.PrintList()
		}
		return
	}

	result, err := formula.Calc()
	if err != nil {
		fmt.Println(err)
		formula.PrintList()
		return
	}
	fmt.Println(result)
}
