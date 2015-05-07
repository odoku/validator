Package of form value validators for Go.


## Installation

Use go get.

```
go get github.com/odoku/validator
```

And import packagae into your code.

```go
import (
    "github.com/odoku/validator"
)
```


## Usage

### Create validator from struct

```go
type Book struct {
    Title string `form:"title" valid:"required(), max_length(50)"`
    Price int    `form:"price" valid:"int(), min(0)"`
}

v := validator.NewWith(Book{})
```

### Create validator manually

```go
rules := validator.Rules{
    "title": validator.Functions{
        validator.NewFunction(validator.IsRequired),
        validator.NewFunction(validator.MaxLength, 50),
    },
    "price": validator.Functions{
        validator.NewFunction(validator.IsInt),
        validator.NewFunction(validator.MinValue, 0),
    },
}

v := validator.New(rules)
```

### Add validation rules

```go
type Book struct {
    Title string `form:"title" valid:"required(), max_length(50)"`
    Price int    `form:"price" valid:"int(), min(0)"`
}

v := validator.NewWith(Book{})

v.AddRule(
    'title',
    validator.NewFunction(validator.MinLength, 3),
)
```

### Create custom validation

```go
f := func(value string) error {
    if value != "foo" {
        return validator.ValidationError("This field must be 'foo'.")
    }
}

v.AddRule('title', validator.Function{Func: f})
```

or like this

```go
f := func(value string) error {
    if value != "foo" {
        return validator.ValidationError("This field must be 'foo'.")
    }
}

err := validator.SetFunc("foo", f)
if err != nil {
    // Invalid function.
    fmt.Printf("%v: %v", err, f)
}

type Book struct {
    Title string `form:"title" valid:"required(), foo()"`
    Price int    `form:"price" valid:"int(), min(0)"`
}

```

### Acceptable function type

```go
func (value string) error

func (value string, args ...interface{}) error

func (form url.Values, value string, args ...interface{}) error

func (values []string) error

func (values []string, args ...interface{}) error

func (form url.Values, values []string, args ...interface{}) error
```

### Builtin functions

```go
// tag: `valid:"required()"`
IsRequired(values []string)

// tag: `valid:"int()"`
IsInt(values []string)

// tag: `valid:"float()"`
IsFloat(values []string)

// tag: `valid:"bool()"`
IsBool(values []string)

// tag: `valid:"date('2006-01-02')"`
IsDate(values []string, format string)

// tag: `valid:"match(/^[a-z]+$/)"`
Match(values []string, pattern *regexp.Regexp)

// tag: `valid:"alpha()"`
IsAlphabet(values []string)

// tag: `valid:"alphanumeric()"`
IsAlphaNumeric(values []string)

// tag: `valid:"email()"`
IsEmail(values []string)

// tag: `valid:"url()"`
IsURL(values []string)

// tag: `valid:"min_length(1)"`
MinLength(values []string, length int)

// tag: `valid:"max_length(100)"`
MaxLength(values []string, length int)

// tag: `valid:"length(1, 100)"`
Length(values []string, min, max)

// tag: `valid:"min(1)"`
MinValue(values []string, value int)

// tag: `valid:"max(100)"`
MaxValue(values []string, value int)

// tag: `valid:"between(1, 100)"`
Between(values []string, min, max)

// tag: `valid:"contains('foo', 'bar', 'buz', 'qux')"`
Contains(values []string, list ...interface{})

// tag: `valid:"if(field, true, required())"`
Condition(form url.Values, values []string, field string, condition interface{}, function Function)
```
