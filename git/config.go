package git

import (
	"os"
)

// Config allows to manipulate the local Git configuration.
type Config struct {
	CommitTemplatePath string
	CommitTemplate     *CommitTemplate
}

// NewConfig returns a new config object based on an existing commit template.
func NewConfig(commitTemplatePath string) (*Config, error) {
	commitTemplate := NewCommitTemplate()
	original, err := os.Open(commitTemplatePath)
	if err != nil {
		return nil, err
	}
	commitTemplate.ReadFrom(original)
	config := &Config{
		commitTemplatePath,
		commitTemplate,
	}
	return config, nil
}
