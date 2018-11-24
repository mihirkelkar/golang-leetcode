package main

import (
	"fmt"
	"math"
)

/*
Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.

Return the quotient after dividing dividend by divisor.

The integer division should truncate toward zero.
*/

func divideTwoIntegers(dividend int, divisor int) int {
	if math.Abs(float64(divisor)) > math.Abs(float64(dividend)) {
		return 0
	}
	var negflag = false
	if divisor*dividend < 0 {
		negflag = true
	}

	quotient := 0
	var remainder = dividend
	for math.Abs(float64(remainder)) >= math.Abs(float64(divisor)) {
		if negflag {
			remainder = remainder + divisor
		} else {
			remainder = remainder - divisor
		}
		quotient++
	}
	if divisor < 0 {
		quotient = quotient * -1
	}
	if dividend < 0 {
		quotient = quotient * -1
	}
	return quotient

}

func main() {
	fmt.Println(divideTwoIntegers(10, 3))
	fmt.Println(divideTwoIntegers(-9, 3))
	fmt.Println(divideTwoIntegers(-9, -2))
	fmt.Println(divideTwoIntegers(10, -9))
}
