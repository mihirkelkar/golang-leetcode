package main

import "unicode"

type Password struct {
	password        string
	numlowercase    int
	numuppercase    int
	numdigit        int
	numrepeated     int
	uppercaserepeat int
	lowercaserepeat int
	digitrepeat     int
}

//define a function type that does an action on the password, like add delete
//or replace character.
type ActionFunc func() int

//define a function type that analyzes the string and returns (true, int)
//if it fits the criteria and by how much
type CheckerFunc func()

//fits the Action Func function type so we can run it together
func (p *Password) AddChar() int {
	//The job of this function is to add a single character.
	//When adding a character, we try to double down on a corner case.
	//so our preference is to add a character that is missing from the criteria.
	var changecount int
	if p.numuppercase == 0 {
		changecount++
		p.numuppercase++
		return changecount
	}

	if p.numlowercase == 0 {
		changecount++
		p.numlowercase++
		return changecount
	}

	if p.numdigit == 0 {
		changecount++
		p.numdigit++
		return changecount
	}

	//default case where all criteria are fulfilled but we're short on length
	changecount++
	p.numlowercase++
	return changecount
}

//Fits the Action Func function type
//Deletes a single character from the string.
//No other logic is involved here accept, we try to delete things.
//that overfill the criteria.
func (p *Password) DeleteChar() int {
	var changecount int
	if p.numlowercase > 1 {
		p.numlowercase--
		changecount++

	} else if p.numuppercase > 1 {
		p.numuppercase--
		changecount++
	} else if p.numdigit > 1 {
		p.numdigit--
		changecount++
	}
	return changecount
}

//Replaces a character with another character.
//Fits the Action Func function type.
//Tries to delete things, and add things.
func (p *Password) ReplaceChar() int {
	var changecount int
	//delete a character, if the changecount returned is non zero,
	//go ahead and add a character
	changecount = p.DeleteChar()
	if changecount > 0 {
		changecount = p.AddChar()
	}
	return changecount
}

//Fits the checker function type.
func (p *Password) CountLowerCase() {
	for _, char := range []rune(p.password) {
		if unicode.IsLower(char) {
			p.numlowercase++
		}
	}
}

func (p *Password) CountUpperCase() {
	for _, char := range []rune(p.password) {
		if unicode.IsUpper(char) {
			p.numuppercase++
		}
	}
}

func (p *Password) CountDigits() {
	for _, char := range []rune(p.password) {
		if unicode.IsDigit(char) {
			p.numuppercase++
		}
	}
}

//write a function that stores repeat char indexes, len, type of char

//if your length is less than 6
//look for repeat chars, we assume that we add chars in between those repeats.
//so we add chars.
//and we subtract repeats.

//if your length is more than 20
//we delete chars from repeats, and we assume we delete
//things that don't introduce more repeats.
//we invlidate repeats.

//if your length is just right between 6 and 20, we replace repeats.
//then fit the other criteria.
