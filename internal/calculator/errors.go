package calculator

import "errors"

var (
	ErrEmptyExpression           = errors.New("empty expression")
	ErrInvalidCharacter          = errors.New("invalid character in expression")
	ErrDivisionByZero            = errors.New("division by zero")
	ErrUnexpectedEnd             = errors.New("unexpected end of expression")
	ErrMissingClosingParenthesis = errors.New("missing closing parenthesis")
	ErrInvalidNumber             = errors.New("invalid number")
)

func IsValidationError(err error) bool {
	return errors.Is(err, ErrEmptyExpression) || errors.Is(err, ErrInvalidCharacter) || errors.Is(err, ErrDivisionByZero) || errors.Is(err, ErrUnexpectedEnd) || errors.Is(err, ErrMissingClosingParenthesis) || errors.Is(err, ErrInvalidNumber)
}
