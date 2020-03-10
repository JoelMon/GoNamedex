// gonamexdex test suite
package gonamedex

import (
	"fmt"
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
	for _, test := range convertToNumberTestCases {
		got, err := convertToNumber(test.input)

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

func TestRemoveDup(t *testing.T) {
	for _, test := range removeDupTestCases {
		got, err := removeDup(test.input)

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

func TestRemoveFirstNumber(t *testing.T) {
	for _, test := range removeFirstNumberTestCases {
		got, err := removeFirstNumber(test.input)

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

func TestRemoveDroppedLetters(t *testing.T) {
	for _, test := range removeDroppedLettersTestCases {
		got, err := removeDroppedLetters(test.input)

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

func TestStripPad(t *testing.T) {
	for _, test := range stripPadTestCases {
		got, err := stripPad(test.input)

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

func ExampleCreate() {
	name := "Joel"
	soundex, _ := Create(name)
	fmt.Println(soundex)
	// Output: J400
}
