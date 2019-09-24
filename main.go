package main

import (
	"fmt"
)

func main() {
	input := "(ADD(SUBTRACT 1 2))"
	tokens := Tokenizer(input)
	p := Parser(tokens)
	fmt.Println(p.Body)
}
