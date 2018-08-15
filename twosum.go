package main

import "fmt"

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
*/

func twosum(arr []int, target int) {
	numMap := make(map[int]int)
	not_found := 0
	for index, value := range arr {
		// check if the target - current num is actually present
		if val, ok := numMap[target-value]; ok {

			fmt.Printf("The indexes to make the value %d are : %d and %d\n", target, val, index)
			not_found++
		}
		numMap[value] = index
	}
	if not_found == 0 {
		fmt.Println("No values that add to the target sum could be found")
	}
}

func main() {
	arr := []int{1, 2, 4, 5}
	twosum(arr, 6)
}
