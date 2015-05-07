package validator

import (
	"errors"
	"net/url"
)

type (
	Function struct {
		Func interface{}
		Args []interface{}
	}

	Functions []*Function
)

var (
	InvalidFunction error = errors.New("Invalid function.")
)

func isValidFunction(f interface{}) bool {
	switch f.(type) {
	case
		func(string) error,
		func(string, ...interface{}) error,
		func(url.Values, string, ...interface{}) error,
		func([]string) error,
		func([]string, ...interface{}) error,
		func(url.Values, []string, ...interface{}) error:
		return true
	default:
		return false
	}
}

func NewFunction(f interface{}, args ...interface{}) *Function {
	if !isValidFunction(f) {
		panic(InvalidFunction)
	}
	return &Function{Func: f, Args: args}
}

func (f *Function) Call(form url.Values, values []string) error {
	switch fn := f.Func.(type) {
	case func(string) error:
		return fn(values[0])
	case func(string, ...interface{}) error:
		return fn(values[0], f.Args...)
	case func(url.Values, string, ...interface{}) error:
		return fn(form, values[0], f.Args...)
	case func([]string) error:
		return fn(values)
	case func([]string, ...interface{}) error:
		return fn(values, f.Args...)
	case func(url.Values, []string, ...interface{}) error:
		return fn(form, values, f.Args...)
	default:
		return nil
	}
}
