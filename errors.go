// Copyright 2024 Jelly Terra
// Use of this source code form is governed under the MIT license.

package httpform

type MissingRequiredFieldError struct {
	Field string
}

func (e MissingRequiredFieldError) Error() string {
	return "missing field: " + e.Field
}

type UnsupportedTypeError struct {
	Field string
}

func (e UnsupportedTypeError) Error() string {
	return "unsupported type: " + e.Field
}
