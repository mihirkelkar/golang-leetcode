package main

import "fmt"

//Sum : Sums an integer slice
func Sum(arr []int) int {
	var count int
	for _, val := range arr {
		count += val
	}
	return count
}

func candyAllocation(ratings []int) int {
	//return with one candy if this is the only element.
	if len(ratings) == 1 {
		return 1
	}
	candyTracker := make([]int, len(ratings))
	for index, value := range ratings {
		//if this is the last element in the array.
		//check if its rating is *greater* than the previous.
		//if so add a candy to it.
		if index == len(ratings)-1 {
			if value > ratings[index-1] {
				candyTracker[index]++
			}
		}

		//if this is the first element in the array.
		//check if its rating is *greater* than the next element.
		//if so add a candy to it
		if index == 0 {
			if value > ratings[index+1] {
				candyTracker[index]++
			}
		}

		//if this is neither the first nor the last element.
		//check which of your neighbors has the higher rating.
		//you need to have more candy than that neighbor if your
		//rating is higher than that neighbor.
		if (index > 0) && (index < len(ratings)-1) {
			beforeRating := ratings[index-1]
			afterRating := ratings[index+1]

			if afterRating > beforeRating {
				if value > afterRating {
					candyTracker[index]++
				}
			} else {
				if value > beforeRating {
					candyTracker[index]++
				}
			}

		}
	}
	return len(ratings) + Sum(candyTracker)

}

func main() {
	fmt.Println(candyAllocation([]int{2, 1, 2}))
	fmt.Println(candyAllocation([]int{1, 2, 2}))
}
