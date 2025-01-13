package pathmux

import (
	"testing"
)

func TestParameterizedRequest_ParameterByName(t *testing.T) {

	tests := []struct{
		ParameterNames []string
		Parameters []string
		Name string
		ExpectedFound bool
		ExpectedValue string
	}{
		{

		},



		{
			ParameterNames: []string{},
			Parameters:     []string{},
			Name: "",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{},
			Parameters:     []string{},
			Name: "no",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{},
			Parameters:     []string{},
			Name: "apple",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{},
			Parameters:     []string{},
			Name: "APPLE",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{},
			Parameters:     []string{},
			Name: "Banana",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{},
			Parameters:     []string{},
			Name: "banana",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{},
			Parameters:     []string{},
			Name: "CHERRY",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{},
			Parameters:     []string{},
			Name: "cherry",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{},
			Parameters:     []string{},
			Name: "dAtE",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{},
			Parameters:     []string{},
			Name: "date",
			ExpectedFound: false,
			ExpectedValue: "",
		},



		{
			ParameterNames: []string{"apple"},
			Parameters:     []string{"once"},
			Name: "",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple"},
			Parameters:     []string{"once"},
			Name: "no",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple"},
			Parameters:     []string{"once"},
			Name: "apple",
			ExpectedFound: true,
			ExpectedValue: "once",
		},
		{
			ParameterNames: []string{"apple"},
			Parameters:     []string{"once"},
			Name: "APPLE",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple"},
			Parameters:     []string{"once"},
			Name: "Banana",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple"},
			Parameters:     []string{"once"},
			Name: "banana",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple"},
			Parameters:     []string{"once"},
			Name: "CHERRY",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple"},
			Parameters:     []string{"once"},
			Name: "cherry",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple"},
			Parameters:     []string{"once"},
			Name: "dAtE",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple"},
			Parameters:     []string{"once"},
			Name: "date",
			ExpectedFound: false,
			ExpectedValue: "",
		},



		{
			ParameterNames: []string{"apple", "Banana"},
			Parameters:     []string{"once",  "Twice"},
			Name: "",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple", "Banana"},
			Parameters:     []string{"once",  "Twice"},
			Name: "no",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple", "Banana"},
			Parameters:     []string{"once",  "Twice"},
			Name: "apple",
			ExpectedFound: true,
			ExpectedValue: "once",
		},
		{
			ParameterNames: []string{"apple", "Banana"},
			Parameters:     []string{"once",  "Twice"},
			Name: "APPLE",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple", "Banana"},
			Parameters:     []string{"once",  "Twice"},
			Name: "Banana",
			ExpectedFound: true,
			ExpectedValue: "Twice",
		},
		{
			ParameterNames: []string{"apple", "Banana"},
			Parameters:     []string{"once",  "Twice"},
			Name: "banana",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple", "Banana"},
			Parameters:     []string{"once",  "Twice"},
			Name: "CHERRY",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple", "Banana"},
			Parameters:     []string{"once",  "Twice"},
			Name: "cherry",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple", "Banana"},
			Parameters:     []string{"once",  "Twice"},
			Name: "dAtE",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple", "Banana"},
			Parameters:     []string{"once",  "Twice"},
			Name: "date",
			ExpectedFound: false,
			ExpectedValue: "",
		},



		{
			ParameterNames: []string{"apple", "Banana", "CHERRY"},
			Parameters:     []string{"once",  "Twice",  "THRICE"},
			Name: "",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple", "Banana", "CHERRY"},
			Parameters:     []string{"once",  "Twice",  "THRICE"},
			Name: "no",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple", "Banana", "CHERRY"},
			Parameters:     []string{"once",  "Twice",  "THRICE"},
			Name: "apple",
			ExpectedFound: true,
			ExpectedValue: "once",
		},
		{
			ParameterNames: []string{"apple", "Banana", "CHERRY"},
			Parameters:     []string{"once",  "Twice",  "THRICE"},
			Name: "APPLE",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple", "Banana", "CHERRY"},
			Parameters:     []string{"once",  "Twice",  "THRICE"},
			Name: "Banana",
			ExpectedFound: true,
			ExpectedValue: "Twice",
		},
		{
			ParameterNames: []string{"apple", "Banana", "CHERRY"},
			Parameters:     []string{"once",  "Twice",  "THRICE"},
			Name: "banana",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple", "Banana", "CHERRY"},
			Parameters:     []string{"once",  "Twice",  "THRICE"},
			Name: "CHERRY",
			ExpectedFound: true,
			ExpectedValue: "THRICE",
		},
		{
			ParameterNames: []string{"apple", "Banana", "CHERRY"},
			Parameters:     []string{"once",  "Twice",  "THRICE"},
			Name: "cherry",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple", "Banana", "CHERRY"},
			Parameters:     []string{"once",  "Twice",  "THRICE"},
			Name: "dAtE",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple", "Banana", "CHERRY"},
			Parameters:     []string{"once",  "Twice",  "THRICE"},
			Name: "date",
			ExpectedFound: false,
			ExpectedValue: "",
		},



		{
			ParameterNames: []string{"apple", "Banana", "CHERRY", "dAtE"},
			Parameters:     []string{"once",  "Twice",  "THRICE", "fOuRcE"},
			Name: "",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple", "Banana", "CHERRY", "dAtE"},
			Parameters:     []string{"once",  "Twice",  "THRICE", "fOuRcE"},
			Name: "no",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple", "Banana", "CHERRY", "dAtE"},
			Parameters:     []string{"once",  "Twice",  "THRICE", "fOuRcE"},
			Name: "apple",
			ExpectedFound: true,
			ExpectedValue: "once",
		},
		{
			ParameterNames: []string{"apple", "Banana", "CHERRY", "dAtE"},
			Parameters:     []string{"once",  "Twice",  "THRICE", "fOuRcE"},
			Name: "APPLE",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple", "Banana", "CHERRY", "dAtE"},
			Parameters:     []string{"once",  "Twice",  "THRICE", "fOuRcE"},
			Name: "Banana",
			ExpectedFound: true,
			ExpectedValue: "Twice",
		},
		{
			ParameterNames: []string{"apple", "Banana", "CHERRY", "dAtE"},
			Parameters:     []string{"once",  "Twice",  "THRICE", "fOuRcE"},
			Name: "banana",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple", "Banana", "CHERRY", "dAtE"},
			Parameters:     []string{"once",  "Twice",  "THRICE", "fOuRcE"},
			Name: "CHERRY",
			ExpectedFound: true,
			ExpectedValue: "THRICE",
		},
		{
			ParameterNames: []string{"apple", "Banana", "CHERRY", "dAtE"},
			Parameters:     []string{"once",  "Twice",  "THRICE", "fOuRcE"},
			Name: "cherry",
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			ParameterNames: []string{"apple", "Banana", "CHERRY", "dAtE"},
			Parameters:     []string{"once",  "Twice",  "THRICE", "fOuRcE"},
			Name: "dAtE",
			ExpectedFound: true,
			ExpectedValue: "fOuRcE",
		},
		{
			ParameterNames: []string{"apple", "Banana", "CHERRY", "dAtE"},
			Parameters:     []string{"once",  "Twice",  "THRICE", "fOuRcE"},
			Name: "date",
			ExpectedFound: false,
			ExpectedValue: "",
		},
	}

	for testNumber, test := range tests {

		var parameterizedRequest ParameterizedRequest
		parameterizedRequest.parameterNames = append([]string(nil), test.ParameterNames...)
		parameterizedRequest.parameters = append([]string(nil), test.Parameters...)

		actualValue, actualFound := parameterizedRequest.ParameterByName(test.Name)

		{
			actual   :=        actualFound
			expected := test.ExpectedFound

			if expected != actual {
				t.Errorf("For test #%d, the actual parameter-by-name-found is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("PARAMETER-NAMES: %#v", test.ParameterNames)
				t.Logf("PARAMETERS:      %#v", test.Parameters)
				t.Logf("NAME: %q", test.Name)
				continue
			}
		}

		{
			actual   :=        actualValue
			expected := test.ExpectedValue

			if expected != actual {
				t.Errorf("For test #%d, the actual parameter-by-name-value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("PARAMETER-NAMES: %#v", test.ParameterNames)
				t.Logf("PARAMETERS: %#v", test.Parameters)
				t.Logf("NAME: %q", test.Name)
				continue
			}
		}
	}
}
