package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field can not be balnk")
		}
	}
}

func (f *Form) MaxLength(field string, maxlen int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > maxlen {
		f.Errors.Add(field, fmt.Sprintf("This field is too long (maximum size is %d characters)", maxlen))
	}
}

func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}
	for _, opt := range opts {
		if value == opt {
			return
		}
	}

	optForErrorStr := ""
	for _, opt := range opts {
		if optForErrorStr == "" {
			optForErrorStr = fmt.Sprintf("\"%s\"", opt)
		} else {
			optForErrorStr = fmt.Sprintf("%s, \"%s\"", optForErrorStr, opt)
		}
	}
	f.Errors.Add(field, fmt.Sprintf("This field value is invalid (valid options are: %s)", optForErrorStr))
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
