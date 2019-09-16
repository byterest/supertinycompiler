package main

import (
	"fmt"
)

func main() {
	input := "ADD 1 2 ADD 1 2 3 4 5 6 7 8 9 0 123132 133123123"
	token := Tokenizer(input)
	fmt.Println(token)
}
