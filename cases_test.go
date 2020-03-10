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

var ConvertToNumberTestCases = []struct {
	description   string
	input         string
	expected      bool
	errorExpected bool
}{
	{
		description: "ConvertToNumber() - letterToReplace1 - b",
		input:       "b",
		expected:    true,
	},
	{
		description: "ConvertToNumber() - letterToReplace1 - f",
		input:       "f",
		expected:    true,
	},
	{
		description: "ConvertToNumber() - letterToReplace1 - p",
		input:       "p",
		expected:    true,
	},
	{
		description: "ConvertToNumber() - letterToReplace1 - !j",
		input:       "j",
		expected:    false,
	},
}
