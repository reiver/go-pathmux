package pathmux

import (
	"testing"
)

func TestParameterizedRequest_ParameterByIndex(t *testing.T) {

	tests := []struct{
		Parameters []string
		Index int
		ExpectedFound bool
		ExpectedValue string
	}{
		{

		},



		{
			Parameters: []string{},
			Index: -2,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{},
			Index: -1,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{},
			Index: 0,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{},
			Index: 1,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{},
			Index: 2,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{},
			Index: 3,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{},
			Index: 4,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{},
			Index: 5,
			ExpectedFound: false,
			ExpectedValue: "",
		},



		{
			Parameters: []string{"once"},
			Index: -2,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{"once"},
			Index: -1,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{"once"},
			Index: 0,
			ExpectedFound: true,
			ExpectedValue: "once",
		},
		{
			Parameters: []string{"once"},
			Index: 1,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{"once"},
			Index: 2,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{"once"},
			Index: 3,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{"once"},
			Index: 4,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{"once"},
			Index: 5,
			ExpectedFound: false,
			ExpectedValue: "",
		},



		{
			Parameters: []string{"once","Twice"},
			Index: -2,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{"once","Twice"},
			Index: -1,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{"once","Twice"},
			Index: 0,
			ExpectedFound: true,
			ExpectedValue: "once",
		},
		{
			Parameters: []string{"once","Twice"},
			Index: 1,
			ExpectedFound: true,
			ExpectedValue: "Twice",
		},
		{
			Parameters: []string{"once","Twice"},
			Index: 2,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{"once","Twice"},
			Index: 3,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{"once","Twice"},
			Index: 4,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{"once","Twice"},
			Index: 5,
			ExpectedFound: false,
			ExpectedValue: "",
		},



		{
			Parameters: []string{"once","Twice","THRICE"},
			Index: -2,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{"once","Twice","THRICE"},
			Index: -1,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{"once","Twice","THRICE"},
			Index: 0,
			ExpectedFound: true,
			ExpectedValue: "once",
		},
		{
			Parameters: []string{"once","Twice","THRICE"},
			Index: 1,
			ExpectedFound: true,
			ExpectedValue: "Twice",
		},
		{
			Parameters: []string{"once","Twice","THRICE"},
			Index: 2,
			ExpectedFound: true,
			ExpectedValue: "THRICE",
		},
		{
			Parameters: []string{"once","Twice","THRICE"},
			Index: 3,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{"once","Twice","THRICE"},
			Index: 4,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{"once","Twice","THRICE"},
			Index: 5,
			ExpectedFound: false,
			ExpectedValue: "",
		},



		{
			Parameters: []string{"once","Twice","THRICE","fOuRcE"},
			Index: -2,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{"once","Twice","THRICE","fOuRcE"},
			Index: -1,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{"once","Twice","THRICE","fOuRcE"},
			Index: 0,
			ExpectedFound: true,
			ExpectedValue: "once",
		},
		{
			Parameters: []string{"once","Twice","THRICE","fOuRcE"},
			Index: 1,
			ExpectedFound: true,
			ExpectedValue: "Twice",
		},
		{
			Parameters: []string{"once","Twice","THRICE","fOuRcE"},
			Index: 2,
			ExpectedFound: true,
			ExpectedValue: "THRICE",
		},
		{
			Parameters: []string{"once","Twice","THRICE","fOuRcE"},
			Index: 3,
			ExpectedFound: true,
			ExpectedValue: "fOuRcE",
		},
		{
			Parameters: []string{"once","Twice","THRICE","fOuRcE"},
			Index: 4,
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Parameters: []string{"once","Twice","THRICE","fOuRcE"},
			Index: 5,
			ExpectedFound: false,
			ExpectedValue: "",
		},
	}

	for testNumber, test := range tests {

		var parameterizedRequest ParameterizedRequest
		parameterizedRequest.parameters = append([]string(nil), test.Parameters...)

		actualValue, actualFound := parameterizedRequest.ParameterByIndex(test.Index)

		{
			actual   :=        actualFound
			expected := test.ExpectedFound

			if expected != actual {
				t.Errorf("For test #%d, the actual parameter-by-index-found is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("PARAMETERS: %#v", test.Parameters)
				t.Logf("INDEX: %d", test.Index)
				continue
			}
		}

		{
			actual   :=        actualValue
			expected := test.ExpectedValue

			if expected != actual {
				t.Errorf("For test #%d, the actual parameter-by-index-value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("PARAMETERS: %#v", test.Parameters)
				t.Logf("INDEX: %d", test.Index)
				continue
			}
		}
	}
}
