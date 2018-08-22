package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

/*
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, two is written as II in Roman numeral,
just two one's added together. Twelve is written as, XII,
which is simply X + II. The number twenty seven is written as XXVII,
which is XX + V + II.
*/
var romanNumeralDict = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func convert_int_to_roman(num int) string {
	keys := make([]int, 0)
	for key, _ := range romanNumeralDict {
		keys = append(keys, key)
	}

	// sort the slice to be largest key first
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	//fmt.Println(keys)
	var buffer bytes.Buffer
	//start with the largest keys
	for i := 0; i < len(keys); {
		//fmt.Printf("This is at index %d\n", i)
		mod_res := num % keys[i]
		if mod_res == num {
			i += 1
		} else {
			//figure out how many multiples that number was
			mults := num / keys[i]
			//fmt.Printf("The mults are %d and the mod_res is %d\n", mults, mod_res)
			buffer.WriteString(strings.Repeat(romanNumeralDict[keys[i]], mults))
			//fmt.Println(buffer.String())
			num = mod_res
			i += 1
		}
		//fmt.Println("************************")
	}
	return buffer.String()
}

func main() {
	fmt.Println(convert_int_to_roman(1239))
}
