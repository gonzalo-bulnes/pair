package template

import (
	"bytes"
	"io"
)

// T is the interface implemented by all templates
type T interface {
	io.ReaderFrom
	io.WriterTo
	String() string
}

// Generic is a generic template, typically a template used for Git commit messages.
type Generic struct {
	content bytes.Buffer
}

// WriteTo implements the io.WriterTo interface.
func (t *Generic) WriteTo(w io.Writer) (n int64, err error) {
	n, err = t.content.WriteTo(w)
	return
}

// ReadFrom implements the io.ReaderFrom interface.
func (t *Generic) ReadFrom(r io.Reader) (n int64, err error) {
	n, err = t.content.ReadFrom(r)
	return
}

// String returns the content of the template as a string.
func (t *Generic) String() string {
	return t.content.String()
}

// New returns a new template.
func New() *Generic {
	return &Generic{}
}
