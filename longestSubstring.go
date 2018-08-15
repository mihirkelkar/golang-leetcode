package main

import "fmt"

/*
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func longestCommonString(str string) int {
	var cur_length int = 1
	var max_length int = 1
	var start_index int = 0
	var cur_index int = 1
	indexmap := make(map[string]int)
	//convert a given string to a rune
	stringrune := []rune(str)
	//fmt.Println(len(stringrune))
	// iterate through the given rune and create an index map
	/*
		for index, value := range stringrune {
			if val, ok := indexmap[string(value)]; ok {
				indexmap[string(value)] = append(val, index)
			} else {
				indexmap[string(value)] = []int{index}
			}
		}
	*/
	indexmap[string(stringrune[0])] = 0
	for cur_index <= len(stringrune)-1 {
		if val, ok := indexmap[string(stringrune[cur_index])]; ok {
			fmt.Println(indexmap)
			fmt.Printf("Found a match for character %s at index %d and the current index is %d\n", string(stringrune[cur_index]), val, cur_index)
			// if the repeated occurence occured before the current substring
			if val < start_index {
				// only index the current element if you're moving ahead
				indexmap[string(stringrune[cur_index])] = cur_index
				cur_index += 1
				cur_length += 1
				if cur_length > max_length {
					max_length = cur_length
				}
			} else {
				//if the repeat occurence occured within the current substring
				start_index = val + 1
				fmt.Printf("Start Index is now %d\n", start_index)
				indexmap[string(stringrune[cur_index])] = cur_index
				cur_index += 1
			}
		} else {
			//only index the current element if you're moving ahead
			indexmap[string(stringrune[cur_index])] = cur_index
			cur_index += 1
			cur_length += 1
			if cur_length > max_length {
				max_length = cur_length
			}

		}

	}
	fmt.Println(max_length)
	return 0
}

func main() {
	longestCommonString("pwwkew")
}
