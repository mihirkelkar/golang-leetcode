package main

import (
	"strconv"
)

type Input int

func (ip Input) NumberComp() int {
	//convert the number
	bnum := strconv.FormatInt(int64(ip), 2)
	var newstring string
	for _, num := range bnum {
		if string(num) == "0" {
			newstring += string(num)
		} else if string(num) == "1" {
			newstring += string(num)
		}
	}
	//fmt.Println(newstring)
	i, err := strconv.ParseInt(newstring, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}
