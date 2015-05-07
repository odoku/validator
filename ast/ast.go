package ast

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/odoku/validator/token"
)

type (
	Function struct {
		Name string
		Args []interface{}
	}
)

func NewInteger(value interface{}) (int, error) {
	return strconv.Atoi(string(value.(*token.Token).Lit))
}

func NewFloat(value interface{}) (float64, error) {
	return strconv.ParseFloat(string(value.(*token.Token).Lit), 64)
}

func NewString(value interface{}, quote string) (string, error) {
	s := string(value.(*token.Token).Lit)

	if quote != "" {
		s = strings.Trim(s, quote)
	}

	return s, nil
}

func NewBoolean(value interface{}) (bool, error) {
	return strconv.ParseBool(string(value.(*token.Token).Lit))
}

func NewRegexp(value interface{}) (*regexp.Regexp, error) {
	return regexp.Compile(strings.Trim(string(value.(*token.Token).Lit), "/"))
}

func NewFunctionList(function interface{}) ([]Function, error) {
	return []Function{function.(Function)}, nil
}

func NewFunction(name interface{}, args interface{}) (Function, error) {
	if args == nil {
		return Function{
			Name: string(name.(*token.Token).Lit),
		}, nil
	} else {
		return Function{
			Name: string(name.(*token.Token).Lit),
			Args: args.([]interface{}),
		}, nil
	}
}

func AppendFunction(list, function interface{}) ([]Function, error) {
	return append(list.([]Function), function.(Function)), nil
}

func NewArgumentList(arg interface{}) ([]interface{}, error) {
	return []interface{}{arg}, nil
}

func AppendArgument(list, arg interface{}) ([]interface{}, error) {
	return append(list.([]interface{}), arg), nil
}
