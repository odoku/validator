package validator

import (
	"net/url"
	"testing"
)

type TestStruct1 struct {
	Hoge int `form:"foo" valid:"required(), int()"`
}

type TestStruct2 struct {
	Piyo bool `form:"bar" valid:"required()"`
	Hoge int  `form:"foo" valid:"int()"`
}

func TestNew(t *testing.T) {
	rule := Rules{
		"title": Functions{
			NewFunction(IsRequired),
			NewFunction(MaxLength, 50),
		},
		"price": Functions{
			NewFunction(IsInt),
			NewFunction(MinValue, 0),
		},
	}
	if rule == nil {
		t.Fail()
	}
}

func TestValidateSuccess(t *testing.T) {
	validator := NewWith(TestStruct1{})

	form := url.Values{"foo": []string{"10"}}
	result, err := validator.Validate(form)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if !result.IsValid {
		t.Fatalf("'%v' is valid values.", form)
	}
	if length := len(result.Errors); length != 0 {
		t.Fatalf("The resutl has %d errors.", length)
	}
}

func TestValidateFailed(t *testing.T) {
	validator := NewWith(TestStruct1{})

	form := url.Values{"foo": []string{"foofoo"}}
	result, err := validator.Validate(form)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if result.IsValid {
		t.Fatalf("'%v' is invalid values.", form)
	}
	if length := len(result.Errors); length == 0 {
		t.Fatalf("The resutl has no errors.")
	}
}

func TestValidateFailedWithEmptyValues(t *testing.T) {
	validator := NewWith(TestStruct1{})

	form := url.Values{}
	result, err := validator.Validate(form)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if result.IsValid {
		t.Fatalf("'%v' is invalid values.", form)
	}
	if length := len(result.Errors); length == 0 {
		t.Fatalf("The resutl has no errors.")
	}
}

func TestValidateSuccessWithMultiValues(t *testing.T) {
	validator := New(Rules{
		"foo": Functions{
			NewFunction(func(values []string) error {
				if values[0] == "10" && values[1] == "20" {
					return nil
				}
				return &ValidationError{Message: "Invalid values."}
			}),
		},
	})

	form := url.Values{
		"foo": []string{"10", "20"},
	}
	result, err := validator.Validate(form)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if !result.IsValid {
		t.Fatalf("Validation Failed. '%v'", result.Errors)
	}
	if length := len(result.Errors); length != 0 {
		t.Fatalf("The resutl has %d errors.", length)
	}
}

func TestCustomeValidation(t *testing.T) {
	validator := NewWith(TestStruct1{})

	validator.AddRule("foo", NewFunction(func(value string) error {
		if value != "10" {
			return &ValidationError{Message: "The field value is not '10'"}
		}
		return nil
	}))

	form := url.Values{"foo": []string{"10"}}
	result, err := validator.Validate(form)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if !result.IsValid {
		t.Fatalf("Validation Failed. '%v'", result.Errors)
	}
	if length := len(result.Errors); length != 0 {
		t.Fatalf("The resutl has %d errors.", length)
	}
}

func TestDynamicValidation(t *testing.T) {
	validator := NewWith(TestStruct2{})

	validator.AddRule("foo", NewFunction(func(form url.Values, values []string, args ...interface{}) error {
		if form.Get("bar") == "true" {
			if err := IsRequired(values); err != nil {
				return err
			}
		}
		return nil
	}))

	form := url.Values{
		"bar": []string{"true"},
		"foo": []string{""},
	}
	result, err := validator.Validate(form)
	if err != nil {
		t.Fatalf("%v", err)
	}

	if result.IsValid {
		t.Fatalf("'%v' is invalid values.", form)
	}
	if length := len(result.Errors); length != 1 {
		t.Fatalf("The resutl has no errors or more errors.")
	}
}
