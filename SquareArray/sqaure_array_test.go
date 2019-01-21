package main

import (
	"testing"
)

func TestSquareArray(t *testing.T) {
	var ip ipArray = []int{-4, -1, 0, 3, 10}
	op := ip.SquareElements()
	if !((op[0] == 0) && (op[1] == 1) && (op[2] == 9) && (op[3] == 16) && (op[4] == 100)) {
		t.Errorf("This should have returned 0, 1, 9, 16, 100")
	}

}
