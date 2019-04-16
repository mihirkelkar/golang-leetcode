package main

import "fmt"

func strstr(pattern string, substring string) int {
	if len(pattern) == 0 {
		return 0
	}
	for pos, char := range pattern {
		if string(char) == string(substring[0]) {
			counter := 1
			for ; counter < len(substring); counter++ {
				if pattern[pos+counter] == substring[counter] {
				}
			}
			if counter == len(substring) {
				return pos
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(strstr("hello", "ll"))
	fmt.Println(strstr("aaaaa", "bba"))
}
