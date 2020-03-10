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
	letterToDrop     = []string{"a", "e", "i", "o", "u", "y", "h", "w"}
	letterToReplace1 = []string{"b", "f", "p", "v"}
	letterToReplace2 = []string{"c", "g", "j", "k", "q", "s", "x", "z"}
	letterToReplace3 = []string{"d", "t"}
	letterToReplace4 = []string{"l"}
	letterToReplace5 = []string{"m", "n"}
	letterToReplace6 = []string{"r"}
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

// convertToNumber takes a name and converts each letter to its
// coresponding numerical value.
func convertToNumber(name string) (string, error) {

	// count is the length of name
	var count int

	// find the length of the string name
	for range name {
		count++
	}

	soundex := make([]string, count)

	// append to soundex the coresponding number for each consentient and also
	// letters that does not have a matching number. ie vowels, h, and w
	for _, letter := range name {
		if contains(letterToReplace1, string(letter)) {
			soundex = append(soundex, "1")
		} else if contains(letterToReplace2, string(letter)) {
			soundex = append(soundex, "2")
		} else if contains(letterToReplace3, string(letter)) {
			soundex = append(soundex, "3")
		} else if contains(letterToReplace4, string(letter)) {
			soundex = append(soundex, "4")
		} else if contains(letterToReplace5, string(letter)) {
			soundex = append(soundex, "5")
		} else if contains(letterToReplace6, string(letter)) {
			soundex = append(soundex, "6")
		} else {
			soundex = append(soundex, string(letter))
		}
	}

	// converts soundex back to a string
	name = strings.Join(soundex, "")

	return name, nil
}

// contains checks to see if a letter is contained within a string slice.
func contains(slice []string, letter string) bool {

	for _, r := range slice {
		if letter == r {
			return true
		}
	}
	return false
}

// removeDup removes numbers that are duplicated following the algorithm
// set by the soundex.
func removeDup(name string) (string, error) {

	rName := []rune(name)
	// lastLetter holds the last letter compared in the for loop.
	lastLetter := rName[0]

	for i, letter := range rName[1:] {
		if letter == lastLetter {
			rName[i] = '?'
			lastLetter = letter
		} else if (letter == 'h' || letter == 'w') && (i+2 < len(rName)) {
			if rName[i] == rName[i+2] {
				rName[i] = '?'
				lastLetter = letter
			}

		} else {
			lastLetter = letter
		}
	}

	sName := string(rName)
	sName = strings.Replace(sName, "?", "", -1)
	return sName, nil
}
