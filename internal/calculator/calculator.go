package calculator

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

const epsilon = 1e-9

func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")

	if expression == "" {
		return 0, ErrEmptyExpression
	}

	tokens, err := tokenize(expression)
	if err != nil {
		return 0, err
	}

	return parseExpression(&tokens)
}

func isZero(x float64) bool {
	return x > -epsilon && x < epsilon
}

func tokenize(expression string) ([]string, error) {
	var tokens []string
	var current strings.Builder

	for _, r := range expression {
		if unicode.IsDigit(r) || r == '.' {
			current.WriteRune(r)
		} else if strings.ContainsRune("+-*/()", r) {
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
			tokens = append(tokens, string(r))
		} else {
			return nil, ErrInvalidCharacter
		}
	}

	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}

	return tokens, nil
}

func parseExpression(tokens *[]string) (float64, error) {
	result, err := parseTerm(tokens)
	if err != nil {
		return 0, err
	}

	for len(*tokens) > 0 {
		op := (*tokens)[0]
		if op != "+" && op != "-" {
			break
		}
		*tokens = (*tokens)[1:]

		nextTerm, err := parseTerm(tokens)
		if err != nil {
			return 0, err
		}

		if op == "+" {
			result += nextTerm
		} else {
			result -= nextTerm
		}
	}

	return result, nil
}

func parseTerm(tokens *[]string) (float64, error) {
	result, err := parseFactor(tokens)
	if err != nil {
		return 0, err
	}

	for len(*tokens) > 0 {
		op := (*tokens)[0]
		if op != "*" && op != "/" {
			break
		}
		*tokens = (*tokens)[1:]

		nextFactor, err := parseFactor(tokens)
		if err != nil {
			return 0, err
		}

		if op == "*" {
			result *= nextFactor
		} else {
			if isZero(math.Abs(nextFactor)) {
				return 0, ErrDivisionByZero
			}
			result /= nextFactor
		}
	}

	return result, nil
}

func parseFactor(tokens *[]string) (float64, error) {
	if len(*tokens) == 0 {
		return 0, ErrUnexpectedEnd
	}

	token := (*tokens)[0]
	*tokens = (*tokens)[1:]

	if token == "(" {
		result, err := parseExpression(tokens)
		if err != nil {
			return 0, err
		}

		if len(*tokens) == 0 || (*tokens)[0] != ")" {
			return 0, ErrMissingClosingParenthesis
		}
		*tokens = (*tokens)[1:]
		return result, nil
	}

	value, err := strconv.ParseFloat(token, 64)
	if err != nil {
		return 0, ErrInvalidNumber
	}

	return value, nil
}
