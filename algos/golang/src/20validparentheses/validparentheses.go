package validparentheses

import (
	"log"
)

func isValid(s string) bool {
	log.Println("Given s:", s)
	stack := []byte{}
	for p := range s[:] {
		switch s[p] {
		case '(', '{', '[':
			stack = append(stack, s[p])
			log.Printf("Stack after pushing (%v): %+v", s[p], stack)
			break
		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}
			cp := stack[len(stack)-1]
			if (s[p] == ')' && cp != '(') ||
				(s[p] == '}' && cp != '{') ||
				(s[p] == ']' && cp != '[') {
				return false
			}
			stack = stack[0 : len(stack)-1]
			log.Printf("Stack after popping (%v) to check %v: %+v", cp, s[p], stack)
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
