package main

import (
	"fmt"
	"pack1/calc"
	"pack1/other"
)

func main() {
	sum := calc.Add(1, 2)
	fmt.Println(sum)

	other.Test()

}
