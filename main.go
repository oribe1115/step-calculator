package main

import (
	"fmt"

	"github.com/oribe1115/step-calculator/lib"
)

func main() {

	test, r, err := lib.ReadNumber("15.4+10")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(test)
	fmt.Println(r)

	plus, r := lib.ReadPlus(r)

	fmt.Println(plus)
	fmt.Println(r)
}
