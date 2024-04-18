// Copyright 2024 Jelly Terra
// Use of this source code form is governed under the MIT license.

package httpform

import (
	"net/http"
	"net/url"
	"strconv"
)

type Wrapper struct {
	Form url.Values

	Vars   map[string]any
	Values map[string]any
}

// Without default value.

func (w *Wrapper) StringVarRequired(p *string, name string) { w.Vars[name] = p }
func (w *Wrapper) UintVarRequired(p *uint, name string)     { w.Vars[name] = p }
func (w *Wrapper) IntVarRequired(p *int, name string)       { w.Vars[name] = p }

func (w *Wrapper) StringRequired(name string) *string {
	p := new(string)
	w.StringVarRequired(p, name)
	return p
}

func (w *Wrapper) UintRequired(name string) *uint {
	p := new(uint)
	w.UintVarRequired(p, name)
	return p
}

func (w *Wrapper) IntRequired(name string) *int {
	p := new(int)
	w.IntVarRequired(p, name)
	return p
}

// With default value.

func (w *Wrapper) StringVar(p *string, name string, value string) {
	w.Vars[name] = p
	w.Values[name] = value
}

func (w *Wrapper) UintVar(p *uint, name string, value uint) {
	w.Vars[name] = p
	w.Values[name] = value
}

func (w *Wrapper) IntVar(p *int, name string, value int) {
	w.Vars[name] = p
	w.Values[name] = value
}

func (w *Wrapper) String(name string, value string) *string {
	p := new(string)
	w.StringVar(p, name, value)
	return p
}

func (w *Wrapper) Uint(name string, value uint) *uint {
	p := new(uint)
	w.UintVar(p, name, value)
	return p
}

func (w *Wrapper) Int(name string, value int) *int {
	p := new(int)
	w.IntVar(p, name, value)
	return p
}

// Parse

func (w *Wrapper) Parse() error {
	for key, ptr := range w.Vars {
		if w.Form.Has(key) {
			// Value exists.
			literal := w.Form.Get(key)
			switch ptr := ptr.(type) {
			case *int:
				v, err := strconv.ParseInt(literal, 10, 64)
				if err != nil {
					return err
				}
				*ptr = int(v)
			case *uint:
				p, err := strconv.ParseUint(literal, 10, 64)
				if err != nil {
					return err
				}
				*ptr = uint(p)
			case *string:
				*ptr = literal
			default:
				return UnsupportedTypeError{Field: key}
			}
		} else {
			// Value is missing, set as default.
			if v, ok := w.Values[key]; ok {
				switch ptr := ptr.(type) {
				case *int:
					*ptr = v.(int)
				case *uint:
					*ptr = v.(uint)
				case *string:
					*ptr = v.(string)
				default:
					return UnsupportedTypeError{Field: key}
				}
			} else {
				return MissingRequiredFieldError{Field: key}
			}
		}
	}

	return nil
}

func Wrap(values url.Values) Wrapper {
	return Wrapper{
		Form:   values,
		Vars:   map[string]any{},
		Values: map[string]any{},
	}
}

func WrapFromRequest(r *http.Request) (Wrapper, error) {
	err := r.ParseForm()
	if err != nil {
		return Wrapper{}, err
	}

	return Wrap(r.Form), nil
}
