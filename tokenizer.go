package main

import (
	"unicode"
)

// Token the key and value
type Token struct {
	Key   string
	Value string
}

// Tokenizer implementation
func Tokenizer(input string) []Token {
	var current = 0
	var tokens []Token
	L := len(input)
	for current < L {
		char := input[current]
		if char == '(' {
			token := Token{Key: "paren", Value: "("}
			tokens = append(tokens, token)
			current++
			continue
		}
		if char == ')' {
			token := Token{Key: "paren", Value: ")"}
			tokens = append(tokens, token)
			current++
			continue
		}
		if unicode.IsSpace(rune(char)) {
			current++
			continue
		}

		if unicode.IsDigit(rune(char)) {
			value := ""
			for unicode.IsDigit(rune(char)) {
				value = string(char) + value
				current++
				if current >= L {
					break
				}
				char = input[current]
			}
			token := Token{Key: "number", Value: value}
			tokens = append(tokens, token)
			continue
		}

		if unicode.IsLetter(rune(char)) {
			value := ""
			for unicode.IsLetter(rune(char)) {
				value = value + string(char)
				current++
				if current >= L {
					break
				}
				char = input[current]
			}
			token := Token{Key: "name", Value: value}

			tokens = append(tokens, token)
			continue
		}
	}
	return tokens
}
