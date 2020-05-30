package main

import (
	"fmt"

	"github.com/oribe1115/step-calculator/lib"
)

func main() {

	test, err := lib.InitFormula("15.4+10")
	test.PrintList()
	if err != nil {
		fmt.Println(err)
		return
	}

	result := test.Calc()
	fmt.Println(result)
}
