package validator

type (
	ValidationError struct {
		Field   string
		Message string
	}

	ValidationErrors map[string][]*ValidationError
)

func (e *ValidationError) Error() string {
	return e.Message
}

func (e ValidationErrors) AsMap() map[string][]string {
	data := map[string][]string{}
	for field, errors := range e {
		data[field] = make([]string, len(errors))
		for i, err := range errors {
			data[field][i] = err.Message
		}
	}
	return data
}
