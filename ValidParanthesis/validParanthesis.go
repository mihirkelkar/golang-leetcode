package main

import "fmt"

func matchParans(input rune, stack []rune) (bool, []rune) {
	var val rune
	if len(stack) > 0 {
		if input == '}' {
			val = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if val == '{' {
				return true, stack
			}
		}
		if input == ']' {
			val = stack[len(stack)-1]
			if val == '[' {
				stack = stack[:len(stack)-1]
				return true, stack
			}
		}
		if input == ')' {
			val = stack[len(stack)-1]
			if val == '(' {
				stack = stack[:len(stack)-1]
				return true, stack
			}
		}
	}
	return false, stack
}

func detectValidParanthesis(input string) bool {
	var paranstack []rune
	var ok bool
	for _, chr := range input {
		if chr == '{' || chr == '[' || chr == '(' {
			paranstack = append(paranstack, chr)
		} else {
			if ok, paranstack = matchParans(chr, paranstack); !ok {
				fmt.Println("This is an invalid paranthesis string")
				return false
			}
		}
	}
	if len(paranstack) > 0 {
		fmt.Println("This is an invalid paranthesis string")
		return false
	}
	fmt.Println("This is a valid paranthesis string")
	return true
}

func main() {

}
