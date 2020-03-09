// Package gonamedex is implementation of the soundex algorithm as mentioned in
// https://www.wikiwand.com/en/Soundex.
package gonamedex

import (
	"errors"
	"strings"
	"unicode"
)

var letterToDrop = []rune{'a', 'e', 'i', 'o', 'u', 'y', 'h', 'w'}
var letterToReplace1 = []rune{'b', 'f', 'p', 'v'}
var letterToReplace2 = []rune{'c', 'g', 'j', 'k', 'q', 's', 'x', 'z'}
var letterToReplace3 = []rune{'d', 't'}
var letterToReplace4 = []rune{'l'}
var letterToReplace5 = []rune{'m', 'n'}
var letterToReplace6 = []rune{'r'}

// makeLower takes a string and performs formatting on the string before returning the string
// with all charactors lowercased.
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
