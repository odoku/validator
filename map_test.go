package validator

import "testing"

func TestGetFunc(t *testing.T) {
	_, err := GetFunc("int")
	if err != nil {
		t.Fatalf("Function 'int' is declared.")
	}

	_, err = GetFunc("unknown")
	if err == nil {
		t.Fatalf("Function 'unknown' is not declared.")
	}
}

func TestSetFunc(t *testing.T) {
	f := func(string) error {
		return nil
	}

	err := SetFunc("foo", f)
	if err != nil {
		t.Fatalf("'%v' is valid function.")
	}

	if _, ok := funcMap["foo"]; !ok {
		t.Fatalf("Function 'foo' is declared")
	}
}

func TestGetErrorMessage(t *testing.T) {
	_, err := GetErrorMessage("int")
	if err != nil {
		t.Fatalf("%v", err)
	}

	_, err = GetErrorMessage("unknown")
	if err == nil {
		t.Fatalf("Error message for 'unknown' is not declared.")
	}
}

func TestSetErrorMessage(t *testing.T) {
	m := "Error Message"
	SetErrorMessage("foo", m)

	if m != errMsg["foo"] {
		t.Fatalf("Error message for 'foo' is not declared")
	}
}
