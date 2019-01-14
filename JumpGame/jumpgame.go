package main

/*
Given an array of non-negative integers, you are initially
positioned at the first index of the array.
Each element in the array represents your maximum jump length at that position.
 Determine if you are able to reach the last index.
For example: A = [2,3,1,1,4], return true. A = [3,2,1,0,4], return false.


Solution : At any given point, your reach is determined by i + A[i].
Lets keep track of your max reach.
If at anypoint your max reach >= len(array) return true
if at any point, you are at a place where i + A[i] <= max && a[0]
which means that there is no path forward, return false.
*/

func jumpGame(arr []int) bool {
	var max int
	for index, value := range arr {
		if index+value >= len(arr)-1 {
			return true
		}

		if index+value > max {
			max = index + value
		}

		if index+value <= max && value == 0 {
			return false
		}
	}
	return true
}
