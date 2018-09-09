package git

import (
	"os"
)

// Config allows to manipulate the local Git configuration.
type Config struct {
	CommitTemplatePath string
	CommitTemplate     *CommitTemplate
}

// Apply writes the commit template to disk.
func (c *Config) Apply() error {
	_ = os.Remove(c.CommitTemplatePath)
	template, err := os.OpenFile(c.CommitTemplatePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, err = c.CommitTemplate.WriteTo(template)
	if err != nil {
		return err
	}
	return nil
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
