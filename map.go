package validator

import (
	"fmt"
)

var (
	funcMap map[string]interface{}
	errMsg  map[string]string
)

func init() {
	funcMap = map[string]interface{}{
		"required":     IsRequired,
		"int":          IsInt,
		"float":        IsFloat,
		"bool":         IsBool,
		"date":         IsDate,
		"match":        Match,
		"alpha":        IsAlphabet,
		"alphanumeric": IsAlphaNumeric,
		"email":        IsEmail,
		"url":          IsURL,
		"min_length":   MinLength,
		"max_length":   MaxLength,
		"length":       Length,
		"min":          MinValue,
		"max":          MaxValue,
		"between":      Between,
		"contains":     Contains,
		"if":           Condition,
	}

	errMsg = map[string]string{
		"required":     "This field is required.",
		"int":          "This field must be a integer.",
		"float":        "This field must be a float.",
		"bool":         "This field must be a bool.",
		"date":         "Enter a valid date.",
		"match":        "Enter a valid value.",
		"alpha":        "Enter a valid email address.",
		"alphanumeric": "This field must be alphabet.",
		"email":        "This field must be alphanumeric.",
		"url":          "Enter a valid url.",
		"min_length":   "Require more than %d charactors.",
		"max_length":   "Require %d charactors or less.",
		"length":       "This field must be at least %d characters %d characters or less.",
		"min":          "This field must be %d or more.",
		"max":          "This field must be %d or less.",
		"between":      "This field must be at %d or more and %d or less.",
		"contains":     "Value %s is not a valid choice.",
	}
}

func GetFunc(name string) (interface{}, error) {
	f, ok := funcMap[name]
	if !ok {
		return nil, fmt.Errorf("Undefined function name. '%s'", name)
	}
	return f, nil
}

func SetFunc(name string, function interface{}) error {
	if !isValidFunction(function) {
		return InvalidFunction
	}
	funcMap[name] = function
	return nil
}

func GetErrorMessage(name string, args ...interface{}) (string, error) {
	m, ok := errMsg[name]
	if !ok {
		return "", fmt.Errorf("Error message for '%s' is not found.", name)
	}
	if len(args) > 0 {
		m = fmt.Sprintf(m, args...)
	}
	return m, nil
}

func SetErrorMessage(name, message string) {
	errMsg[name] = message
}
