package main

import "fmt"

/*
Given two sorted arrays, Write a funtion that accepts both of these
arrays a single combined returns another sorted array
*/
func combineSortedArray(arrOne []int, arrTwo []int) []int {
	indexOne := 0
	indexTwo := 0
	var combArray []int
	for (indexOne < len(arrOne)) && (indexTwo < len(arrTwo)) {
		if arrOne[indexOne] <= arrTwo[indexTwo] {
			combArray = append(combArray, arrOne[indexOne])
			indexOne++
		} else {
			combArray = append(combArray, arrTwo[indexTwo])
			indexTwo++
		}
	}

	for indexOne < len(arrOne) {
		combArray = append(combArray, arrOne[indexOne])
		indexOne++
	}

	for indexTwo < len(arrTwo) {
		combArray = append(combArray, arrTwo[indexTwo])
		indexTwo++
	}

	return combArray

}

func main() {
	arrOne := []int{1, 5, 8, 19, 22}
	arrTwo := []int{0, 4, 11, 23, 49, 1000, 100001}
	fmt.Println(combineSortedArray(arrOne, arrTwo))
}
