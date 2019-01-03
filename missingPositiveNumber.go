package main

import (
	"fmt"
	"sort"
)

/*
Missing Positive Number
Given an unsorted integer array, find the smallest missing positive integer.

Example 1:

Input: [1,2,0]
Output: 3
Example 2:

Input: [3,4,-1,1]
Output: 2
Example 3:

Input: [7,8,9,11,12]
Output: 1
Note:

Your algorithm should run in O(n) time and uses constant extra space.

1. Sort the array.
2. Start with 0 and binary search each number until you have one you can't find
*/

func binarySearch(start int, end int, search int, array []int) bool {
	//fmt.Println("search is ", search)
	for start <= end {
		middle := (start + end) / 2

		//fmt.Println("end is ", end)
		//fmt.Println("-- -- -- -- --")
		if array[middle] == search {
			return true
		}
		if array[middle] < search {
			start = middle + 1
		}
		if array[middle] > search {
			end = middle - 1
		}
	}
	return false
}

func findMissingPosInt(array []int) {
	sort.Ints(array)
	var x int
	for ; binarySearch(0, len(array)-1, x, array); x++ {
	}
	fmt.Println("missing number is ", x)
}

func main() {
	findMissingPosInt([]int{1, 2, 0})
	findMissingPosInt([]int{3, 4, -1, 1})
}
