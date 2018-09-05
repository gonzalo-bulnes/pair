package git

import (
	"regexp"

	"github.com/gonzalo-bulnes/pair/github"
	"github.com/gonzalo-bulnes/pair/template"
)

// CommitTemplate is a template used for Git commit messages.
type CommitTemplate struct {
	*template.Generic
}

// CoAuthor returns the co-author if any. The second return value indicates
// whether or not a co-author was mentioned in the template.
func (t *CommitTemplate) CoAuthor() (string, bool) {
	re := regexp.MustCompile(github.CoAuthorRegexp)
	found := author(re.FindAllStringSubmatch(t.String(), 1))
	if found == nil {
		return "", false
	}
	return found.coAuthor, true
}

// NewCommitTemplate returns a new Git commit template.
func NewCommitTemplate() *CommitTemplate {
	return &CommitTemplate{
		template.New(),
	}
}

type matched struct {
	line     string
	coAuthor string
}

func author(m [][]string) *matched {
	if m == nil {
		return nil
	}
	return &matched{
		line:     m[0][0],
		coAuthor: m[0][1],
	}
}
