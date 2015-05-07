package validator

import (
	"fmt"
	"net/url"
	"testing"
)

func TestFunctionType1(t *testing.T) {
	f := NewFunction(func(values []string) error {
		for _, v := range values {
			if v != "10" {
				return &ValidationError{Message: "The value is not '10'"}
			}
		}
		return nil
	})

	err := f.Call(url.Values{}, []string{"10"})
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestFunctionType2(t *testing.T) {
	f := NewFunction(func(values []string, args ...interface{}) error {
		for _, v := range values {
			if v != args[0] {
				return &ValidationError{Message: fmt.Sprintf("The value is not '%v'", args[0])}
			}
		}
		return nil
	}, "10")

	err := f.Call(url.Values{}, []string{"10"})
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestFunctionType3(t *testing.T) {
	f := NewFunction(func(form url.Values, values []string, args ...interface{}) error {
		if form.Get("foo") != args[0] {
			return &ValidationError{Message: fmt.Sprintf("The value is not '%v'", args[0])}
		}
		return nil
	}, "10")

	err := f.Call(url.Values{"foo": []string{"10"}}, []string{""})
	if err != nil {
		t.Fatalf("%v", err)
	}
}
