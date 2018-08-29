package main

import "fmt"

func find_max_area_container(arr []int) int {
	var maxarea int = 0
	var lowptr_index int = 0
	var hghptr_index int = len(arr) - 1
	var smallersize int = 0
	var curmaxarea int = 0
	var lowptrsmallflag bool = false
	//The idea here is to start from both ends moving inwards.
	//calculate the area, update the max variable and
	//move the pointer with the smaller height inwards
	for lowptr_index < hghptr_index {
		if arr[lowptr_index] < arr[hghptr_index] {
			smallersize = arr[lowptr_index]
			lowptrsmallflag = true
		} else {
			smallersize = arr[hghptr_index]
			lowptrsmallflag = false
		}
		curmaxarea = smallersize * (hghptr_index - lowptr_index)
		if curmaxarea > maxarea {
			maxarea = curmaxarea
		}

		if lowptrsmallflag {
			lowptr_index++
		} else {
			hghptr_index--
		}

	}
	return maxarea
}

func main() {
	var container_arr = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Printf("The maximum area of water in this container is %d\n", find_max_area_container(container_arr))
}
