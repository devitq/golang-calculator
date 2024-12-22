package calculator_test

import (
	"errors"
	"testing"

	"calc_service/internal/calculator"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
		err        error
	}{
		{"", 0, calculator.ErrEmptyExpression},
		{"2 + 2", 4, nil},
		{"3 * (2 + 4)", 18, nil},
		{"10 / 2", 5, nil},
		{"10 / 0", 0, calculator.ErrDivisionByZero},
		{"(5 + 2", 0, calculator.ErrMissingClosingParenthesis},
		{"2.5 + 2.5", 5, nil},
		{"abc + 1", 0, calculator.ErrInvalidCharacter},
	}

	for _, test := range tests {
		t.Run(test.expression, func(t *testing.T) {
			result, err := calculator.Calc(test.expression)
			if err != test.err {
				t.Errorf("unexpected error: got %v, want %v", err, test.err)
			}
			if result != test.expected {
				t.Errorf("unexpected result: got %v, want %v", result, test.expected)
			}
		})
	}
}

func TestIsValidationError(t *testing.T) {
	tests := []struct {
		err        error
		isExpected bool
	}{
		{calculator.ErrEmptyExpression, true},
		{calculator.ErrInvalidCharacter, true},
		{calculator.ErrDivisionByZero, true},
		{calculator.ErrUnexpectedEnd, true},
		{calculator.ErrMissingClosingParenthesis, true},
		{calculator.ErrInvalidNumber, true},
		{nil, false},
		{errors.New("random error"), false},
	}

	for _, test := range tests {
		testName := "nil"
		if test.err != nil {
			testName = test.err.Error()
		}

		t.Run(testName, func(t *testing.T) {
			isValid := calculator.IsValidationError(test.err)
			if isValid != test.isExpected {
				t.Errorf("unexpected validation result: got %v, want %v", isValid, test.isExpected)
			}
		})
	}
}
