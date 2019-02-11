package main

import (
	"testing"
)

func TestValidParanthesis(t *testing.T) {
	//check for valid paranthesis
	if ok := detectValidParanthesis("(())"); !ok {
		t.Errorf("Error : Expected this to be a valid paranthesis")
	}
}

func TestNestedValidParanthesis(t *testing.T) {
	//check for valid paranthesis
	if ok := detectValidParanthesis("({[[{()}]]})"); !ok {
		t.Errorf("Error : Expected this to be a valid paranthesis")
	}
}

func TestInvalidParanthesis(t *testing.T) {
	//check for invalid paranthesis
	if ok := detectValidParanthesis("(("); ok {
		t.Errorf("Error : Expected this to be an invalid paranthesis")
	}
}

func TestInvalidCharacter(t *testing.T) {
	//check for invalid character in input string
	if ok := detectValidParanthesis("3343"); ok {
		t.Errorf("Error : Expected this to be an invalid paranthesis")
	}
}
