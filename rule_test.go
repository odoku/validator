package validator

import (
	"reflect"
	"testing"
)

func TestGetSubstance(t *testing.T) {
	var i interface{} = &struct{}{}

	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Ptr {
		t.Fatalf("'%v' is not pointer value", v.Kind())
	}

	v = getSubstance(i)
	if v.Kind() == reflect.Ptr {
		t.Fatalf("'%v' is pointer value", v.Kind())
	}
}

func TestCreateRulesFromStruct(t *testing.T) {
	rules, err := createRulesFromStruct(struct {
		Hoge int `form:"foo" valid:"required(), int()"`
	}{})
	if err != nil {
		t.Fatalf("%v", err)
	}

	vfs, ok := rules["foo"]
	if !ok {
		t.Fatalf("Undefined 'foo'")
	}

	if length := len(vfs); length != 2 {
		t.Fatalf("%d is not 2", length)
	}

	callFunctions := []reflect.Value{
		reflect.ValueOf(IsRequired),
		reflect.ValueOf(IsInt),
	}

Parent:
	for _, vf := range vfs {
		fv := reflect.ValueOf(vf.Func)
		for _, x := range callFunctions {
			if fv.Pointer() == x.Pointer() {
				continue Parent
			}
		}
		t.Fatalf("'%v' is invalid function", fv)
	}
}
