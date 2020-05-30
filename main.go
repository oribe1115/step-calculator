package main

import (
	"fmt"

	"github.com/oribe1115/step-calculator/lib"
)

func main() {

	num, r, err := lib.ReadNumber("15.4+10")
	test := lib.CreateNumberToken(num)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(test)
	fmt.Println(r)
}
