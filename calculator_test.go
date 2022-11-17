package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		description string
		a           int
		b           int
		expVal      int
	}{
		{"Simple addition case", 5, 5, 10},
		{"Addition with negative numbers", 5, -10, -5},
	}

	for i, testCase := range tests {
		gotVal := add(testCase.a, testCase.b)

		if gotVal != testCase.expVal {
			t.Errorf("Test(%v):[%v]: Want(%v), Got(%v)", i, testCase.description, testCase.expVal, gotVal)
		}
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		description string
		a           int
		b           int
		expVal      int
	}{
		{"Simple subtraction case", 10, 3, 7},
		{"Sub with negative numbers", 5, -29, 34},
	}

	for i, testCase := range tests {
		gotVal := sub(testCase.a, testCase.b)

		if gotVal != testCase.expVal {
			t.Errorf("Test(%v):[%v]: Exp(%v), Got(%v)", i, testCase.description, testCase.expVal, gotVal)
		}
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		description string
		a           int
		b           int
		expVal      int
	}{
		{"Simple multiplication case", 5, 5, 25},
		{"Mul with negative numbers", -23, 2, -46},
	}

	for i, testCase := range tests {
		gotVal := mul(testCase.a, testCase.b)

		if gotVal != testCase.expVal {
			t.Errorf("Test(%v):[%v]: Exp(%v), Got(%v)", i, testCase.description, testCase.expVal, gotVal)
		}
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		description string
		a           int
		b           int
		expVal      int
		expError    error
	}{
		{"Proper division case", 25, 5, 5, nil},
		{"Improper division case", 25, 3, 8, nil},
		{"Division by zero", 25, 0, 0, myError{"Division by zero occurred"}},
	}

	for i, testCase := range tests {
		gotVal, gotError := div(testCase.a, testCase.b)

		if gotVal != testCase.expVal {
			t.Errorf("Test(%v):[%v]: Exp(%v), Got(%v)", i, testCase.description, testCase.expVal, gotVal)
		}

		if gotError != testCase.expError {
			t.Errorf("Test(%v):[%v]: Exp(%v), Got(%v)", i, testCase.description, testCase.expError, gotError)
		}
	}
}

func TestMod(t *testing.T) {
	tests := []struct {
		description string
		a           int
		b           int
		expVal      int
		expError    error
	}{
		{"Simple modulo operation", 123, 10, 3, nil},
		{"Modulo by negative number", 123, -10, -3, nil},
		{"Modulo by 0 case", 12, 0, 0, myError{"Modulo by zero occurred"}},
	}

	for i, testCase := range tests {
		gotVal, gotError := mod(testCase.a, testCase.b)

		if gotVal != testCase.expVal {
			t.Errorf("Test(%v):[%v]: Exp(%v), Got(%v)", i, testCase.description, testCase.expVal, gotVal)
		}

		if gotError != testCase.expError {
			t.Errorf("Test(%v):[%v]: Exp(%v), Got(%v)", i, testCase.description, testCase.expError, gotError)
		}
	}
}

func TestCalculator(t *testing.T) {
	tests := []struct {
		description string
		a           int
		b           int
		operator    string
		expVal      int
		expError    error
	}{
		{"Simple addition case", 5, 5, "+", 10, nil},
		{"Simple subtraction case", 10, 3, "-", 7, nil},
		{"Simple multiplication case", 5, 5, "*", 25, nil},
		{"Proper division case", 25, 5, "/", 5, nil},
		{"Improper division case", 25, 3, "/", 8, nil},
		{"Division by zero", 25, 0, "/", 0, myError{"Division by zero occurred"}},
		{"Simple modulo operation", 123, 10, "%", 3, nil},
		{"Modulo by negative number", 123, -10, "%", -3, nil},
		{"Modulo by 0 case", 12, 0, "%", 0, myError{"Modulo by zero occurred"}},
		{"Invalid operator", 5, 2, "#", 0, myError{"Invalid operator"}},
	}

	for i, testCase := range tests {
		gotVal, gotError := calculator(testCase.a, testCase.b, testCase.operator)

		if gotVal != testCase.expVal {
			t.Errorf("Test(%v):[%v]: Exp(%v), Got(%v)", i, testCase.description, testCase.expVal, gotVal)
		}

		if gotError != testCase.expError {
			t.Errorf("Test(%v):[%v]: Exp(%v), Got(%v)", i, testCase.description, testCase.expError, gotError)
		}
	}
}

// end of the program
