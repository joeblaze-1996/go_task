package main

import (
	"fmt"
	"strconv"
)

func main() {
	var decimal int64
	fmt.Print("pls enter a decimal number:")
	fmt.Scanln(&decimal)
	output := strconv.FormatInt(decimal, 2)
	fmt.Print("Output ", output)

}
