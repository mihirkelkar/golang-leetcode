package main

import (
	"sort"
)

type ipArray []int

func (ip ipArray) SquareElements() ipArray {
	//checking squares of elements.
	newArray := make(ipArray, 0)
	for _, val := range ip {
		newArray = append(newArray, val*val)
	}
	sort.Ints(newArray)
	return newArray
}
