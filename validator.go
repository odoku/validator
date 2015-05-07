package validator

import "net/url"

type (
	Validator struct {
		Rules Rules
	}

	Result struct {
		IsValid bool
		Errors  ValidationErrors
	}
)

func New(rules Rules) *Validator {
	return &Validator{Rules: rules}
}

func NewWith(s interface{}) *Validator {
	rules, err := createRulesFromStruct(s)
	if err != nil {
		panic(err)
	}
	return New(rules)
}

func (v *Validator) AddRule(field string, rules ...*Function) {
	v.Rules[field] = append(v.Rules[field], rules...)
}

func (v *Validator) Validate(form url.Values) (*Result, error) {
	result := &Result{Errors: ValidationErrors{}}

	for field, fs := range v.Rules {
		values, ok := form[field]
		if !ok {
			values = []string{""}
		}

		for _, f := range fs {
			err := f.Call(form, values)
			if err != nil {
				e, ok := err.(*ValidationError)
				if !ok {
					return nil, err
				}

				e.Field = field
				result.Errors[field] = append(result.Errors[field], e)
			}
		}
	}

	result.IsValid = len(result.Errors) == 0

	return result, nil
}
