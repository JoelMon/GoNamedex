package gonamedex

var makeLowerTestCases = []struct {
	description   string
	input         string
	expected      string
	errorExpected bool
}{
	{
		description: "makeLower() - Lowercased name",
		input:       "joel",
		expected:    "joel",
	},
	{
		description: "makeLower() - Uppercased name",
		input:       "Mary",
		expected:    "mary",
	},
	{
		description: "makeLower() - Mixed cased name",
		input:       "DaViD",
		expected:    "david",
	},
	{
		description: "makeLower() - One char",
		input:       "c",
		expected:    "c",
	},
	{
		description: "makeLower() - One char Uppercased",
		input:       "C",
		expected:    "c",
	},
	{
		description:   "makeLower() - No char",
		input:         "",
		expected:      "error",
		errorExpected: true,
	},
	{
		description:   "makeLower() - Only a space",
		input:         " ",
		expected:      "error",
		errorExpected: true,
	},
	{
		description:   "makeLower() - Space in front of name",
		input:         " Jose",
		expected:      "error",
		errorExpected: true,
	},
	{
		description:   "makeLower() - Space at end of name",
		input:         "Jose ",
		expected:      "error",
		errorExpected: true,
	},
	{
		description:   "makeLower() - More than one name",
		input:         "Robert J. Smith",
		expected:      "error",
		errorExpected: true,
	},
	{
		description:   "makeLower() - Special Charactors-1",
		input:         "+",
		expected:      "error",
		errorExpected: true,
	},
	{
		description:   "makeLower() - Number",
		input:         "6",
		expected:      "error",
		errorExpected: true,
	},
	{
		description:   "makeLower() - Number Mixed",
		input:         "James3",
		expected:      "error",
		errorExpected: true,
	},
	{
		description:   "makeLower() - Special Charactors-3",
		input:         "@@$5",
		expected:      "error",
		errorExpected: true,
	},
	{
		description:   "makeLower() - Special Charactors Mixed",
		input:         "John-Smith",
		expected:      "error",
		errorExpected: true,
	},
}

var convertToNumberTestCases = []struct {
	description   string
	input         string
	expected      string
	errorExpected bool
}{
	{
		description: "convertToNumber() - letterToReplace1 - b",
		input:       "b",
		expected:    "1",
	},
	{
		description: "convertToNumber() - letterToReplace1 - bfp",
		input:       "bfp",
		expected:    "111",
	},
	{
		description: "convertToNumber() - letterToReplace2 - cgjkqsxz",
		input:       "cgjkqsxz",
		expected:    "22222222",
	},
	{
		description: "convertToNumber() - letterToReplace3 - dt",
		input:       "dt",
		expected:    "33",
	},
	{
		description: "convertToNumber() - letterToReplace4 - l",
		input:       "l",
		expected:    "4",
	},
	{
		description: "convertToNumber() - letterToReplace5 - mn",
		input:       "mn",
		expected:    "55",
	},
	{
		description: "convertToNumber() - letterToReplace6 - r",
		input:       "r",
		expected:    "6",
	},
	{
		description: "convertToNumber() - letterToReplace - joel",
		input:       "joel",
		expected:    "2oe4",
	},
	{
		description: "convertToNumber() - letterToReplace - frank",
		input:       "frank",
		expected:    "16a52",
	},
	{
		description: "convertToNumber() - letterToReplace - amy",
		input:       "amy",
		expected:    "a5y",
	},
	{
		description: "convertToNumber() - letterToReplace - leroy",
		input:       "leroy",
		expected:    "4e6oy",
	},
}
