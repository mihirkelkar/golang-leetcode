package main

import "fmt"

/*
Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].
*/

func find_first_last(target int, arr []int) []int {
	range_values := []int{0, 0}
	middle_index := binary_search(target, 0, len(arr)-1, arr)

	if middle_index == -1 {
		fmt.Println("Middle index not found")
		return []int{-1, -1}
	}
	//now that we have a middle index that is not -1,
	//recursively go to the left till the middle index returns a -1
	lowerbound := func(low_index int) int {
		var lowerbound = low_index
		for {
			fmt.Printf("The low index now is %d\n", low_index)
			low_index = binary_search(target, 0, low_index, arr)
			fmt.Printf("The low index now is %d\n", low_index)
			if low_index != -1 {
				lowerbound = low_index
				low_index = low_index - 1
			} else {
				return lowerbound
			}
		}
	}

	upperbound := func(high_index int) int {
		var upperbound = high_index
		for {
			fmt.Printf("The high index now is %d\n", high_index)
			high_index = binary_search(target, high_index, len(arr)-1, arr)
			fmt.Printf("The high index now is %d\n", high_index)
			if high_index != -1 {
				upperbound = high_index
				high_index = high_index + 1
			} else {
				return upperbound
			}
		}
	}
	low_index := middle_index
	high_index := middle_index
	range_values[0], range_values[1] = lowerbound(low_index), upperbound(high_index)
	return range_values
}

func binary_search(target int, low int, high int, arr []int) int {
	for low < high {
		middle := (low + high) / 2
		if arr[middle] == target {
			return middle
		} else {
			if arr[middle] < target {
				low = middle + 1
			} else {
				high = middle - 1
			}
		}
	}
	return -1
}

func main() {
	temp := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(find_first_last(6, temp))
	fmt.Println(find_first_last(8, temp))
}
