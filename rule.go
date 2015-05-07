package validator

import (
	"fmt"
	"reflect"

	"github.com/odoku/validator/ast"
	"github.com/odoku/validator/lexer"
	"github.com/odoku/validator/parser"
)

type (
	Rules map[string]Functions
)

func createRulesFromStruct(s interface{}) (Rules, error) {
	v := getSubstance(s)
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%T Is not struct value.", s)
	}

	rules := Rules{}

	for i := 0; i < v.NumField(); i++ {
		t := v.Type().Field(i)

		tag := t.Tag.Get("valid")
		if tag == "" {
			continue
		}

		field := t.Tag.Get("form")
		if field == "" {
			field = t.Name
		}

		p := parser.NewParser()
		s := lexer.NewLexer([]byte(tag))
		result, err := p.Parse(s)
		if err != nil {
			return nil, fmt.Errorf("Could not parse %s field tag: %v", field, err)
		}

		functions := result.([]ast.Function)
		vfs := make(Functions, len(functions))

		for i, function := range functions {
			f, err := GetFunc(function.Name)
			if err != nil {
				return nil, err
			}
			vfs[i] = &Function{Func: f, Args: function.Args}
		}

		rules[field] = vfs
	}

	return rules, nil
}

func getSubstance(s interface{}) reflect.Value {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}
