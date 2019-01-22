package main

import (
	"testing"
)

func TestNumberComplement(t *testing.T) {
	//ip := Input(2)
	//var ip Input = 2
	//if ip.NumberComp() != 2 {
	//	t.Errorf("The output was supposed to be 2")
	//}
	ip := Input(1)
	//fmt.Println(ip.NumberComp())
	if ip.NumberComp() != 0 {
		t.Errorf("The output was supposed to be 0")
	}
}
