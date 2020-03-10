// gonamexdex test suite
package gonamedex

import (
	"testing"
)

// TestMakeLower is someting
func TestMakeLower(t *testing.T) {
	for _, test := range makeLowerTestCases {
		got, err := makeLower(test.input)

		if test.errorExpected {
			// check if err is of type error
			var _ error = err

			// we expect error
			if err == nil {
				t.Fatalf("FAIL: %s - Expected an error when the input is %q.", test.description, test.input)
			}
		} else {
			if err != nil {
				t.Fatalf("FAIL: %s - An error was not expected with input %q.\n Error msg: %q.", test.description, test.input, err)
			}
		}

		if got != test.expected {
			t.Fatalf("FAIL: %s - Expected %q but got %q with input %q", test.description, test.expected, got, test.input)
		}
		t.Logf("Passed - %s", test.description)
	}
}

func TestConvertToNumber(t *testing.T) {
	for _, test := range ConvertToNumberTestCases {
		got, err := ConvertToNumber(test.input)

		if test.errorExpected {
			// check if err is of type error
			var _ error = err

			// we expect error
			if err == nil {
				t.Fatalf("FAIL: %s - Expected an error when the input is %q.", test.description, test.input)
			}
		} else {
			if err != nil {
				t.Fatalf("FAIL: %s - An error was not expected with input %q.\n Error msg: %q.", test.description, test.input, err)
			}
		}

		if got != test.expected {
			t.Fatalf("FAIL: %s - Expected %v but got %v with input %v", test.description, test.expected, got, test.input)
		}
		t.Logf("Passed - %s", test.description)
	}
}

// func ExampleCreate() {
// 	name, _ := Create("Joel")
// 	fmt.Println(name)
// 	//Output: joel
// }

// func ExampleConvertToNumber() {
// 	name := ConvertToNumber("Joel")
// 	fmt.Println(name)
// 	//Output: 2oe4
// }

// func ExampleContains() {
// 	var listRunes = []rune{'a', 'e', 'o'}

// 	inList := Contains(listRunes, 'a')
// 	fmt.Println(inList)
// 	// Output: true
// }
