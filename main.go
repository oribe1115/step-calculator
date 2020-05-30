package main

import (
	"fmt"

	"github.com/oribe1115/step-calculator/lib"
)

func main() {

	test := lib.InitFormula("15.4+10")
	test.PrintList()

	result := test.Calc()
	fmt.Println(result)
}
