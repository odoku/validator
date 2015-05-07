package validator

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/odoku/validator/ast"
	"github.com/odoku/validator/lexer"
	"github.com/odoku/validator/parser"
)

func TestParse(t *testing.T) {
	tag := `
		foo(
			nil,
			10, +10, -10,
			10.14, +10.14, -10.14,
			true, false,
			foo_bar, HogePiyo,
			'*&^%',
			/[\w]+/
		)
	`
	p := parser.NewParser()
	s := lexer.NewLexer([]byte(tag))

	result, err := p.Parse(s)
	if err != nil {
		t.Fatalf("%v", err)
	}

	values := []ast.Function{
		ast.Function{Name: "foo", Args: []interface{}{
			nil,
			10, 10, -10,
			10.14, 10.14, -10.14,
			true, false,
			"foo_bar", "HogePiyo",
			"*&^%",
			regexp.MustCompile(`[\w]+`),
		}},
	}

	if !reflect.DeepEqual(result, values) {
		t.Fatalf("Is not equals. '%v' and '%v", result, values)
	}
}

func TestParseWith2Values(t *testing.T) {
	tag := "foo(), bar(10)"
	p := parser.NewParser()
	s := lexer.NewLexer([]byte(tag))

	result, err := p.Parse(s)
	if err != nil {
		t.Fatalf("%v", err)
	}

	values := []ast.Function{
		ast.Function{Name: "foo"},
		ast.Function{Name: "bar", Args: []interface{}{10}},
	}

	if !reflect.DeepEqual(result, values) {
		t.Fatalf("Is not equals. '%v' and '%v", result, values)
	}
}

func TestParseWithComplexValues(t *testing.T) {
	tag := "foo(bar, 10, foo(bar, baz, qux)), xxx(10)"
	p := parser.NewParser()
	s := lexer.NewLexer([]byte(tag))

	result, err := p.Parse(s)
	if err != nil {
		t.Fatalf("%v", err)
	}

	values := []ast.Function{
		ast.Function{Name: "foo", Args: []interface{}{
			"bar", 10,
			ast.Function{
				Name: "foo",
				Args: []interface{}{"bar", "baz", "qux"},
			},
		}},
		ast.Function{Name: "xxx", Args: []interface{}{10}},
	}

	if !reflect.DeepEqual(result, values) {
		t.Fatalf("Is not equals. '%v' and '%v", result, values)
	}
}

func TestParseWithSameName(t *testing.T) {
	tag := "foo(10, 20, 30), foo(40, 50, 60)"
	p := parser.NewParser()
	s := lexer.NewLexer([]byte(tag))

	result, err := p.Parse(s)
	if err != nil {
		t.Fatalf("%v", err)
	}

	values := []ast.Function{
		ast.Function{Name: "foo", Args: []interface{}{10, 20, 30}},
		ast.Function{Name: "foo", Args: []interface{}{40, 50, 60}},
	}

	if !reflect.DeepEqual(result, values) {
		t.Fatalf("Is not equals. '%v' and '%v", result, values)
	}
}

func TestParseWithMultiLineValue(t *testing.T) {
	tag := `
		foo(10, 20, 30),
		foo(40, 50, 60)
	`
	p := parser.NewParser()
	s := lexer.NewLexer([]byte(tag))

	result, err := p.Parse(s)
	if err != nil {
		t.Fatalf("%v", err)
	}

	values := []ast.Function{
		ast.Function{Name: "foo", Args: []interface{}{10, 20, 30}},
		ast.Function{Name: "foo", Args: []interface{}{40, 50, 60}},
	}

	if !reflect.DeepEqual(result, values) {
		t.Fatalf("Is not equals. '%v' and '%v", result, values)
	}
}
