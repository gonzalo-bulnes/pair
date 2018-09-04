package template

import (
	"bytes"
	"io"
	"regexp"
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

// CoAuthor returns the co-author if any. The second return value indicates
// whether or not a co-author was mentioned in the template.
func (t *T) CoAuthor() (string, bool) {
	re := regexp.MustCompile(`[Cc]o-[Aa]uthored-[Bb]y: (.[^\n]*)`)
	found := author(re.FindAllStringSubmatch(t.content.String(), 1))
	if found == nil {
		return "", false
	}
	return found.coAuthor, true
}

// New returns a new template.
func New() *T {
	return &T{}
}

type match struct {
	match    string
	coAuthor string
}

func author(m [][]string) *match {
	if m == nil {
		return nil
	}
	return &match{
		match:    m[0][0],
		coAuthor: m[0][1],
	}
}