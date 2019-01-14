package main

import (
	"testing"
)

func TestJumpGame(t *testing.T) {
	//check that this ir returns true for an empty array
	if jumpGame([]int{}) != true {
		t.Errorf("Empty lists shoudl retrurn true")
	}

	if jumpGame([]int{2, 3, 1, 1, 4}) != true {
		t.Errorf("This should have returned a true")
	}

	if jumpGame([]int{3, 2, 1, 0, 4}) != false {
		t.Errorf("This should have returned a false")
	}
}
