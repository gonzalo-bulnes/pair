package template

import (
	"bytes"
	"io"
)

// T is a template, typically a template used for Git commit messages.
type T struct {
	content bytes.Buffer
}

// WriteTo implements the io.WriterTo interface.
func (t *T) WriteTo(w io.Writer) (n int64, err error) {
	n, err = t.content.WriteTo(w)
	return
}

// ReadFrom implements the io.ReaderFrom interface.
func (t *T) ReadFrom(r io.Reader) (n int64, err error) {
	n, err = t.content.ReadFrom(r)
	return
}

// New returns a new template.
func New() *T {
	return &T{}
}
