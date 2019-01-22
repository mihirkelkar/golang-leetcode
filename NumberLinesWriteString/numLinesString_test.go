package main

import "testing"

func TestStringNumLines(t *testing.T) {
	width := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	str := "abcdefghijklmnopqrstuvwxyz"
	ls := InitLineString(width, str)
	a, b := ls.Calculate()
	if a != 3 && b != 60 {
		t.Errorf("Should have returned 3 and 60")
	}
}
