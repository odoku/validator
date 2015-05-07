package validator

import (
	"net/url"
	"regexp"
	"testing"

	"github.com/odoku/validator/ast"
)

func TestIsRequired(t *testing.T) {
	v := []string{"abc"}
	if err := IsRequired(v); err != nil {
		t.Errorf("'%v' is not empty value", v)
	}

	v = []string{""}
	if err := IsRequired(v); err == nil {
		t.Errorf("'%v' is empty value", v)
	}
}

func TestIsInt(t *testing.T) {
	v := []string{"123", "-123", "+123", "0", ""}
	if err := IsInt(v); err != nil {
		t.Errorf("'%v' is integer value.", v)
	}

	v = []string{"123", "abc", "abc123", "123abc"}
	if err := IsInt(v); err == nil {
		t.Errorf("'%v' is not integer value.", v)
	}
}

func TestIsFloat(t *testing.T) {
	v := []string{"0.12", "-0.123", "+12.123", ""}
	if err := IsFloat(v); err != nil {
		t.Errorf("'%v' is float value.", v)
	}

	v = []string{"12.4", "abc", "abc123", "123abc"}
	if err := IsFloat(v); err == nil {
		t.Errorf("'%v' is not float value.", v)
	}
}

func TestIsBool(t *testing.T) {
	v := []string{"true", "false", ""}
	if err := IsBool(v); err != nil {
		t.Errorf("'%v' is bool value.", v)
	}

	v = []string{"true", "abc", "abc123", "123abc"}
	if err := IsBool(v); err == nil {
		t.Errorf("'%v' is not bool value.", v)
	}
}

func TestIsDate(t *testing.T) {
	v := []string{"2015-10-10", "2014-03-02", ""}
	if err := IsDate(v, "2006-01-02"); err != nil {
		t.Errorf("'%v' is date value.", v)
	}

	v = []string{"2015-10-10", "abc", "abc123", "123abc"}
	if err := IsDate(v, "2006-01-02"); err == nil {
		t.Errorf("'%v' is not date value.", v)
	}
}

func TestMatch(t *testing.T) {
	v := []string{"foo10", "bar20", ""}
	if err := Match(v, regexp.MustCompile(`\w{3}\d{2}`)); err != nil {
		t.Errorf("'%v' is matching value.", v)
	}

	v = []string{"ssssss", "abc", "abc123", "123abc"}
	if err := Match(v, regexp.MustCompile(`\w{3}\d{2}`)); err == nil {
		t.Errorf("'%v' is not matching value.", v)
	}
}

func TestIsAlphabet(t *testing.T) {
	v := []string{"foo", "atoz", ""}
	if err := IsAlphabet(v); err != nil {
		t.Errorf("'%v' is alphabet.", v)
	}

	v = []string{"1233", "ahohgoe3234"}
	if err := IsAlphabet(v); err == nil {
		t.Errorf("'%v' is not alphabet.", v)
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	v := []string{"foo", "bar123", "12314", ""}
	if err := IsAlphaNumeric(v); err != nil {
		t.Errorf("'%v' is alphabet.", v)
	}

	v = []string{"hglsk###", "+1203"}
	if err := IsAlphaNumeric(v); err == nil {
		t.Errorf("'%v' is not alphabet.", v)
	}
}

func TestIsEmail(t *testing.T) {
	v := []string{"foo@foo.com", "foo.foo@foo.com", ""}
	if err := IsEmail(v); err != nil {
		t.Errorf("'%v' is email address.", v)
	}

	v = []string{"ssssss", "abc", "abc123", "123abc"}
	if err := IsEmail(v); err == nil {
		t.Errorf("'%v' is not email address.", v)
	}
}

func TestMinLength(t *testing.T) {
	v := []string{"12345", "5文字以上の日本語", ""}
	if err := MinLength(v, 5); err != nil {
		t.Errorf("'%v' is valid value.", v)
	}

	v = []string{"1234", "foo"}
	if err := MinLength(v, 5); err == nil {
		t.Errorf("'%v' is invalid value.", v)
	}
}

func TestMaxLength(t *testing.T) {
	v := []string{"12345", "3文字", ""}
	if err := MaxLength(v, 5); err != nil {
		t.Errorf("'%v' is valid value.", v)
	}

	v = []string{"123456", "laksjlakjdlkajs"}
	if err := MaxLength(v, 5); err == nil {
		t.Errorf("'%v' is invalid value.", v)
	}
}

func TestLength(t *testing.T) {
	v := []string{"12345", "1234567890", ""}
	if err := Length(v, 5, 10); err != nil {
		t.Errorf("'%v' is valid value.", v)
	}

	v = []string{"1234", "12345678900"}
	if err := Length(v, 5, 10); err == nil {
		t.Errorf("'%v' is invalid value.", v)
	}
}

func TestMinValue(t *testing.T) {
	v := []string{"5", "9999", ""}
	if err := MinValue(v, 5); err != nil {
		t.Errorf("'%v' is valid value.", v)
	}

	v = []string{"4", "foo"}
	if err := MinValue(v, 5); err == nil {
		t.Errorf("'%v' is invalid value.", v)
	}
}

func TestMaxValue(t *testing.T) {
	v := []string{"5", "-2", ""}
	if err := MaxValue(v, 5); err != nil {
		t.Errorf("'%v' is valid value.", v)
	}

	v = []string{"6", "9999"}
	if err := MaxValue(v, 5); err == nil {
		t.Errorf("'%v' is invalid value.", v)
	}
}

func TestBetween(t *testing.T) {
	v := []string{"5", "10", "6"}
	if err := Between(v, 5, 10); err != nil {
		t.Errorf("'%v' is valid value.", v)
	}

	v = []string{"4", "11"}
	if err := Between(v, 5, 10); err == nil {
		t.Errorf("'%v' is invalid value.", v)
	}
}

func TestIsURL(t *testing.T) {
	v := []string{"http://google.com", "https://google.com?foo=bar", ""}
	if err := IsURL(v); err != nil {
		t.Errorf("'%v' is url.", v)
	}

	v = []string{"httt://ssssss", "abc", "abc123", "123abc"}
	if err := IsURL(v); err == nil {
		t.Errorf("'%v' is not url.", v)
	}
}

func TestContains(t *testing.T) {
	v := []string{"foo", "bar", ""}
	if err := Contains(v, "foo", "bar"); err != nil {
		t.Errorf("'%v' is valid value.", v)
	}

	v = []string{"baz", "qux"}
	if err := Contains(v, "foo", "bar"); err == nil {
		t.Errorf("'%v' is not date value.", v)
	}
}

func TestCondition(t *testing.T) {
	form := url.Values{
		"foo": []string{"10"},
	}
	values := []string{"1982/10", "1982/01"}

	args := []interface{}{
		"foo",
		10,
		Function{Func: IsDate, Args: []interface{}{"2006/01/02"}},
	}
	err := Condition(form, values, args...)
	if err == nil {
		t.Errorf("Process must be an error.")
	}

	args = []interface{}{
		"foo",
		101010,
		ast.Function{Name: "date", Args: []interface{}{"2006/01/02"}},
	}
	err = Condition(form, values, args...)
	if err != nil {
		t.Errorf("Process must be succeed. '%v'", err)
	}
}
