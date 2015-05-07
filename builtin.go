package validator

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/odoku/validator/ast"
)

var (
	rxEmail        = regexp.MustCompile("^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$")
	rxURL          = regexp.MustCompile(`^((ftp|http|https):\/\/)?(\S+(:\S*)?@)?((([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|((www\.)?)?(([a-z\x{00a1}-\x{ffff}0-9]+-?-?_?)*[a-z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-z\x{00a1}-\x{ffff}]{2,}))?)|localhost)(:(\d{1,5}))?((\/|\?|#)[^\s]*)?$`)
	rxAlphabet     = regexp.MustCompile("^[a-zA-Z]+$")
	rxAlphaNumeric = regexp.MustCompile("^[a-zA-Z0-9]+$")
)

func isEmpty(values []string) bool {
	if len(values) == 0 {
		return true
	}
	for _, v := range values {
		if v != "" {
			return false
		}
	}
	return true
}

func IsRequired(values []string) error {
	if isEmpty(values) {
		m, err := GetErrorMessage("required")
		if err != nil {
			return err
		}
		return &ValidationError{Message: m}
	}

	return nil
}

func IsInt(values []string) error {
	for _, v := range values {
		if v == "" {
			continue
		}
		_, err := strconv.Atoi(v)
		if err != nil {
			m, err := GetErrorMessage("int")
			if err != nil {
				return err
			}
			return &ValidationError{Message: m}
		}
	}
	return nil
}

func IsFloat(values []string) error {
	for _, v := range values {
		if v == "" {
			continue
		}
		_, err := strconv.ParseFloat(v, 64)
		if err != nil {
			m, err := GetErrorMessage("float")
			if err != nil {
				return err
			}
			return &ValidationError{Message: m}
		}
	}
	return nil
}

func IsBool(values []string) error {
	for _, v := range values {
		if v == "" {
			continue
		}
		_, err := strconv.ParseBool(v)
		if err != nil {
			m, err := GetErrorMessage("bool")
			if err != nil {
				return err
			}
			return &ValidationError{Message: m}
		}
	}
	return nil
}

func IsDate(values []string, args ...interface{}) error {
	var format string
	if len(args) == 0 {
		format = "2006-01-02"
	} else {
		format = args[0].(string)
	}

	for _, v := range values {
		if v == "" {
			continue
		}
		_, err := time.Parse(format, v)
		if err != nil {
			m, err := GetErrorMessage("date")
			if err != nil {
				return err
			}
			return &ValidationError{Message: m}
		}
	}

	return nil
}

func Match(values []string, args ...interface{}) error {
	if len(args) < 1 {
		return fmt.Errorf("MinLength argument must be 1 argument.")
	}

	r, ok := args[0].(*regexp.Regexp)
	if !ok {
		return fmt.Errorf("First argument must be Regexp.")
	}

	for _, v := range values {
		if v == "" {
			continue
		}
		if !r.MatchString(v) {
			m, err := GetErrorMessage("match")
			if err != nil {
				return err
			}
			return &ValidationError{Message: m}
		}
	}
	return nil
}

func IsAlphabet(values []string) error {
	for _, v := range values {
		if v == "" {
			continue
		}
		if !rxAlphabet.MatchString(v) {
			m, err := GetErrorMessage("alpha")
			if err != nil {
				return err
			}
			return &ValidationError{Message: m}
		}
	}
	return nil
}

func IsAlphaNumeric(values []string) error {
	for _, v := range values {
		if v == "" {
			continue
		}
		if !rxAlphaNumeric.MatchString(v) {
			m, err := GetErrorMessage("alphanumeric")
			if err != nil {
				return err
			}
			return &ValidationError{Message: m}
		}
	}
	return nil
}

func IsEmail(values []string) error {
	for _, v := range values {
		if v == "" {
			continue
		}
		if !rxEmail.MatchString(v) {
			m, err := GetErrorMessage("email")
			if err != nil {
				return err
			}
			return &ValidationError{Message: m}
		}
	}
	return nil
}

func IsURL(values []string) error {
	for _, v := range values {
		if v == "" {
			continue
		}
		if !rxURL.MatchString(v) {
			m, err := GetErrorMessage("url")
			if err != nil {
				return err
			}
			return &ValidationError{Message: m}
		}
	}
	return nil
}

func MinLength(values []string, args ...interface{}) error {
	if len(args) < 1 {
		return fmt.Errorf("MinLength argument must be 1 argument.")
	}

	length, ok := args[0].(int)
	if !ok {
		return fmt.Errorf("First argument must be integer value.")
	}

	for _, v := range values {
		if v == "" {
			continue
		}
		if utf8.RuneCountInString(v) < length {
			m, err := GetErrorMessage("min_length", length)
			if err != nil {
				return err
			}
			return &ValidationError{Message: m}
		}
	}
	return nil
}

func MaxLength(values []string, args ...interface{}) error {
	if len(args) < 1 {
		return fmt.Errorf("MaxLength argument must be 1 argument.")
	}

	length, ok := args[0].(int)
	if !ok {
		return fmt.Errorf("First argument must be integer value.")
	}

	for _, v := range values {
		if v == "" {
			continue
		}
		if utf8.RuneCountInString(v) > length {
			m, err := GetErrorMessage("max_length", length)
			if err != nil {
				return err
			}
			return &ValidationError{Message: m}
		}
	}
	return nil
}

func Length(values []string, args ...interface{}) error {
	if len(args) < 1 {
		return fmt.Errorf("Length argument must be 2 arguments.")
	}

	min, ok := args[0].(int)
	if !ok {
		return fmt.Errorf("First argument must be integer value.")
	}

	max, ok := args[1].(int)
	if !ok {
		return fmt.Errorf("First argument must be integer value.")
	}

	for _, v := range values {
		if v == "" {
			continue
		}
		l := utf8.RuneCountInString(v)
		if l < min || l > max {
			m, err := GetErrorMessage("langth", min, max)
			if err != nil {
				return err
			}
			return &ValidationError{Message: m}
		}
	}
	return nil
}

func MinValue(values []string, args ...interface{}) error {
	if len(args) < 1 {
		return fmt.Errorf("MinLength argument must be 1 argument.")
	}

	min, ok := args[0].(int)
	if !ok {
		return fmt.Errorf("First argument must be integer value.")
	}

	for _, v := range values {
		if v == "" {
			continue
		}
		iv, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		if iv < min {
			m, err := GetErrorMessage("min", min)
			if err != nil {
				return err
			}
			return &ValidationError{Message: m}
		}
	}
	return nil
}

func MaxValue(values []string, args ...interface{}) error {
	if len(args) < 1 {
		return fmt.Errorf("MinLength argument must be 1 argument.")
	}

	max, ok := args[0].(int)
	if !ok {
		return fmt.Errorf("First argument must be integer value.")
	}

	for _, v := range values {
		if v == "" {
			continue
		}
		iv, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		if iv > max {
			m, err := GetErrorMessage("max", max)
			if err != nil {
				return err
			}
			return &ValidationError{Message: m}
		}
	}
	return nil
}

func Between(values []string, args ...interface{}) error {
	if len(args) < 1 {
		return fmt.Errorf("Length argument must be 2 arguments.")
	}

	min, ok := args[0].(int)
	if !ok {
		return fmt.Errorf("First argument must be integer value.")
	}

	max, ok := args[1].(int)
	if !ok {
		return fmt.Errorf("First argument must be integer value.")
	}

	for _, v := range values {
		if v == "" {
			continue
		}
		iv, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		if iv < min || iv > max {
			m, err := GetErrorMessage("between", min, max)
			if err != nil {
				return err
			}
			return &ValidationError{Message: m}
		}
	}
	return nil
}

func Contains(values []string, args ...interface{}) error {
Parent:
	for _, v := range values {
		if v == "" {
			continue
		}
		for _, arg := range args {
			if v == fmt.Sprintf("%v", arg) {
				continue Parent
			}
		}
		m, err := GetErrorMessage("contains")
		if err != nil {
			return err
		}
		return &ValidationError{Message: m}
	}
	return nil
}

func Condition(form url.Values, values []string, args ...interface{}) error {
	if len(args) != 3 {
		return fmt.Errorf("Condition argument must be 3 arguments.")
	}

	if form.Get(args[0].(string)) != fmt.Sprintf("%v", args[1]) {
		return nil
	}

	vf, ok := args[2].(Function)
	if !ok {
		function := args[2].(ast.Function)
		f, err := GetFunc(function.Name)
		if err != nil {
			return err
		}
		vf = Function{Func: f, Args: function.Args}
	}

	return vf.Call(form, values)
}
