package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func combinationSum(arr []int, target int) []string {
	//sort the numbers in decreasing order.
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	var result map[string]int
	result = make(map[string]int)
	//now start with the first integer
	for index, value := range arr {
		if value > target {
			continue
		} else {
			if value == target {
				result[strconv.Itoa(value)] = 1
				continue
			} else {
				stackIndex := 0
				stack := make([]int, len(arr)-index)
				var currentIndex int
				currentIndex = index + 1
				for currentIndex < len(arr) {
					if sumSlice(stack)+arr[currentIndex] < target {
						stack[stackIndex] = arr[currentIndex]
						stackIndex++
					} else if sumSlice(stack)+arr[currentIndex] == target {
						stack[stackIndex] = arr[currentIndex]
						sort.Ints(stack)
						result[sliceString(stack)] = 1
						stack[stackIndex] = 0
					}
					currentIndex++
				}
			}
		}
	}
	return getKeys(result)
}

func sumSlice(arr []int) int {
	result := 0
	for _, value := range arr {
		result += value
	}
	return result
}

func sliceString(arr []int) string {
	var IDs []string
	for _, i := range arr {
		if i != 0 {
			IDs = append(IDs, strconv.Itoa(i))
		}
	}
	return strings.Join(IDs, ",")
}

func getKeys(mymap map[string]int) []string {
	keys := make([]string, len(mymap))
	for v := range mymap {
		keys = append(keys, v)
	}
	return keys
}

func main() {
	var arr = []int{2, 5, 2, 1, 2}
	fmt.Println(combinationSum(arr, 5))
	arr = []int{10, 1, 2, 7, 6, 1, 5}
	fmt.Println(combinationSum(arr, 8))
}
