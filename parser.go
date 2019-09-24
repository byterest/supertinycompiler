package main

import (
	"fmt"
)

type node struct {
	Type string
	Name string
	Params []node
	Body   []node
}

// Parser to AST
func Parser(tokens []Token) node {
	fmt.Println("begin parse")
	current := 0
	ast := node{
		Type: "Program",
	}
	fmt.Println("sleeep")
	for current < len(tokens) {
	ast.Body = append(ast.Body, walk(tokens, &current))
	fmt.Println("AST", current)
	}
	return ast
}

func walk(tokens []Token, current *int) node{
	fmt.Println("walk")
	fmt.Println(*current)
	token := tokens[*current]
	fmt.Println("f", *current)
	fmt.Println(token.Value)
	if token.Key == "number" {
		*current = 1 +  *current
		fmt.Println("number", token.Value)
		return node {
			Type: "NumberLiteral",
			Name: token.Value,
		}

	}

	if token.Key == "paren" && token.Value == "(" {
		fmt.Println("paren")
		*current = 1 +  *current
		token = tokens[*current]
		ns := node{
			Type: "CallExpression",
			Name: token.Value,
		}
		*current = 1 +  *current
		token := tokens[*current]

		for (token.Key != "paren") || (token.Key=="paren" && token.Value !=")"){
			fmt.Println("node")
			fmt.Println(token.Key)
			ns.Params = append(ns.Params, walk(tokens, current))
			token = tokens[*current]
		}
		*current = 1 + *current
		return ns
	}
	panic("err")
}