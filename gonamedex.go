// Package gonamedex is an implementation of the American soundex algorithm as mentioned in
// https://en.wikipedia.org/wiki/Soundex#American_Soundex.
package gonamedex

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// Various arrays of letters that has to be replaced or dropped
// when creating the soundex.
var (
	letterToDrop     = []rune{'a', 'e', 'i', 'o', 'u', 'y', 'h', 'w'}
	letterToReplace1 = []rune{'b', 'f', 'p', 'v'}
	letterToReplace2 = []rune{'c', 'g', 'j', 'k', 'q', 's', 'x', 'z'}
	letterToReplace3 = []rune{'d', 't'}
	letterToReplace4 = []rune{'l'}
	letterToReplace5 = []rune{'m', 'n'}
	letterToReplace6 = []rune{'r'}
)

// Create takes a name and returns the corresponding soundex
// for the name.
//
//NOTE
//
// Implimentation incomplete.
func Create(name string) (string, error) {

	convert, err := makeLower(name)
	if err != nil {
		fmt.Println(err)
	}
	return convert, nil
}

// makeLower takes a string and first checks to see if it's an empty string,
// it then checks to see if the string only contains alpha charactors,
// lastly, it returns the input string back in all lowercase.
//
// Code:
//  x, _ := makeLower("Joel")
//  fmt.Println(x)
// Output:
//  joel
func makeLower(name string) (string, error) {

	// check for empty string
	if name == "" {
		return "error", errors.New("string must not be empty, \"\"")
	}
	// check for nonalpha charactors
	for _, letter := range name {
		if !unicode.IsLetter(letter) {
			return "error", errors.New("the entered string contain nonalpha charactors")
		}
	}
	// return lowercased string
	return strings.ToLower(name), nil
}

// ConvertToNumber takes a name and converts each letter to its
// coresponding numerical value.
func ConvertToNumber(name string) (bool, error) {

	// count is the length of name
	var count int

	// find the length of the string name
	for range name {
		count++
	}

	// soundex := make([]rune, count)

	// rName is a rune slice made from the string name
	rName := []rune(name)

	for _, letter := range rName {
		if Contains(letterToReplace1, letter) {
			return true, nil
		}
	}

	return false, nil
}

// Contains checks to see if someting is contained.
func Contains(slice []rune, letter rune) bool {

	for _, r := range slice {
		if letter == r {
			return true
		}
	}
	return false

}
