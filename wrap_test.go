// Copyright 2024 Jelly Terra
// Use of this source code form is governed under the MIT license.

package httpform

import (
	"net/url"
	"testing"
)

func TestWrapper_Parse(t *testing.T) {
	var form = url.Values{}

	w := Wrap(form)

	i := w.Int("i", 1)
	s := w.String("s", "str")

	err := w.Parse()
	if err != nil {
		t.Fatal(err)
	}

	if *i != 1 {
		t.Fatal()
	}

	if *s != "str" {
		t.Fatal()
	}
}
