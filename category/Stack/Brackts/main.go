package main

import (
	"fmt"
	"strings"
)

var bracketMap map[string]string

func init() {
	bracketMap = map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
}

type bracketStack struct {
	bracket []string
}

func newBracketStack() *bracketStack {
	return &bracketStack{}
}

func push(val string, stack *bracketStack) {
	stack.bracket = append(stack.bracket, val)
	for checkBracket(stack) {
	}
}

func checkBracket(stack *bracketStack) bool {
	if len(stack.bracket) > 1 {
		if stack.bracket[len(stack.bracket)-1] == bracketMap[stack.bracket[len(stack.bracket)-2]] {
			pop(stack)
			pop(stack)
			return true
		}
	}
	return false
}

func pop(stack *bracketStack) string {
	str := stack.bracket[len(stack.bracket)-1]
	stack.bracket = stack.bracket[:len(stack.bracket)-1]
	return str
}

func checkEmptyStack(stack bracketStack) bool {
	if len(stack.bracket) == 0 {
		return true
	}
	return false
}

func BalancedBrackets(s string) bool {
	var strArr []string
	strArr = strings.Split(s, "")
	stack := newBracketStack()
	for _, val := range strArr {
		if val == "(" || val == ")" || val == "{" || val == "}" || val == "[" || val == "]" {
			push(val, stack)
		}
	}
	if len(stack.bracket) == 0 {
		return true
	}
	return false
}

func main() {
	testCase1 := "([])(){}(())()()"
	fmt.Println(BalancedBrackets(testCase1))
	testCase2 := "(a)"
	fmt.Println(BalancedBrackets(testCase2))
}
