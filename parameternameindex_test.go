package pathmux

import (
	"testing"
)

func TestParameterNameIndex(t *testing.T) {

	tests := []struct{
		ParameterNames []string
		Name string
		ExpectedFound bool
		ExpectedIndex int
	}{
		{

		},



		{
			ParameterNames: []string{"apple"},
			Name: "",
			ExpectedFound: false,
			ExpectedIndex: 0,
		},
		{
			ParameterNames: []string{"apple"},
			Name: "apple",
			ExpectedFound: true,
			ExpectedIndex: 0,
		},
		{
			ParameterNames: []string{"apple"},
			Name: "APPLE",
			ExpectedFound: false,
			ExpectedIndex: 0,
		},
		{
			ParameterNames: []string{"apple"},
			Name: "Banana",
			ExpectedFound: false,
			ExpectedIndex: 0,
		},
		{
			ParameterNames: []string{"apple"},
			Name: "banana",
			ExpectedFound: false,
			ExpectedIndex: 0,
		},
		{
			ParameterNames: []string{"apple"},
			Name: "CHERRY",
			ExpectedFound: false,
			ExpectedIndex: 0,
		},
		{
			ParameterNames: []string{"apple"},
			Name: "cherry",
			ExpectedFound: false,
			ExpectedIndex: 0,
		},
		{
			ParameterNames: []string{"apple"},
			Name: "dAtE",
			ExpectedFound: false,
			ExpectedIndex: 0,
		},
		{
			ParameterNames: []string{"apple"},
			Name: "XYZ",
			ExpectedFound: false,
			ExpectedIndex: 0,
		},



		{
			ParameterNames: []string{"apple","Banana"},
			Name: "",
			ExpectedFound: false,
			ExpectedIndex: 0,
		},
		{
			ParameterNames: []string{"apple","Banana"},
			Name: "apple",
			ExpectedFound: true,
			ExpectedIndex: 0,
		},
		{
			ParameterNames: []string{"apple","Banana"},
			Name: "APPLE",
			ExpectedFound: false,
			ExpectedIndex: 0,
		},
		{
			ParameterNames: []string{"apple","Banana"},
			Name: "Banana",
			ExpectedFound: true,
			ExpectedIndex: 1,
		},
		{
			ParameterNames: []string{"apple","Banana"},
			Name: "banana",
			ExpectedFound: false,
			ExpectedIndex: 0,
		},
		{
			ParameterNames: []string{"apple","Banana"},
			Name: "CHERRY",
			ExpectedFound: false,
			ExpectedIndex: 0,
		},
		{
			ParameterNames: []string{"apple","Banana"},
			Name: "cherry",
			ExpectedFound: false,
			ExpectedIndex: 0,
		},
		{
			ParameterNames: []string{"apple","Banana"},
			Name: "dAtE",
			ExpectedFound: false,
			ExpectedIndex: 0,
		},
		{
			ParameterNames: []string{"apple","Banana"},
			Name: "XYZ",
			ExpectedFound: false,
			ExpectedIndex: 0,
		},



		{
			ParameterNames: []string{"apple","Banana","CHERRY"},
			Name: "",
			ExpectedFound: false,
			ExpectedIndex: 0,
		},
		{
			ParameterNames: []string{"apple","Banana","CHERRY"},
			Name: "apple",
			ExpectedFound: true,
			ExpectedIndex: 0,
		},
		{
			ParameterNames: []string{"apple","Banana","CHERRY"},
			Name: "APPLE",
			ExpectedFound: false,
			ExpectedIndex: 0,
		},
		{
			ParameterNames: []string{"apple","Banana","CHERRY"},
			Name: "Banana",
			ExpectedFound: true,
			ExpectedIndex: 1,
		},
		{
			ParameterNames: []string{"apple","Banana","CHERRY"},
			Name: "banana",
			ExpectedFound: false,
			ExpectedIndex: 0,
		},
		{
			ParameterNames: []string{"apple","Banana","CHERRY"},
			Name: "CHERRY",
			ExpectedFound: true,
			ExpectedIndex: 2,
		},
		{
			ParameterNames: []string{"apple","Banana","CHERRY"},
			Name: "cherry",
			ExpectedFound: false,
			ExpectedIndex: 0,
		},
		{
			ParameterNames: []string{"apple","Banana","CHERRY"},
			Name: "dAtE",
			ExpectedFound: false,
			ExpectedIndex: 0,
		},
		{
			ParameterNames: []string{"apple","Banana","CHERRY"},
			Name: "XYZ",
			ExpectedFound: false,
			ExpectedIndex: 0,
		},
	}

	for testNumber, test := range tests {

		actualIndex, actualFound := parameterNameIndex(test.Name, test.ParameterNames...)

		{
			actual   :=        actualFound
			expected := test.ExpectedFound

			if expected != actual {
				t.Errorf("For test #%d, the actual parameter-name-found is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("PARAMETER-NAME: %#v", test.ParameterNames)
				t.Logf("NAME: %q", test.Name)
				continue
			}
		}

		{
			actual   :=        actualIndex
			expected := test.ExpectedIndex

			if expected != actual {
				t.Errorf("For test #%d, the actual parameter-name-index is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("PARAMETER-NAME: %#v", test.ParameterNames)
				t.Logf("NAME: %q", test.Name)
				continue
			}
		}
	}
}
