// Package git provides primitives to manage the Git commit.template configuration.
package git

import (
	"bytes"
	"os/exec"
	"regexp"
)

const globalConfigOutputRegexp = `file:.git/config`
const commitTemplatePathRegexp = `config\s+(.*)\s*`

// NoCommitTemplateConfigurationError indicates that commit.template is unset.
type NoCommitTemplateConfigurationError struct{}

// Error implements the error interface.
func (e *NoCommitTemplateConfigurationError) Error() string {
	return "No commit.template configuration found"
}

// CLI allows to operate the system's Git command-line interface.
type CLI struct {
	getCommitTemplatePath func() (out bytes.Buffer, err error)
	setCommitTemplate     func(path string) (err error)
	unsetCommitTemplate   func() (err error)
}

// NewCLI returns a new CLI that wraps the Git command-line interface
func NewCLI() *CLI {
	return &CLI{
		getCommitTemplatePath: _getCommitTemplatePath,
		setCommitTemplate:     _setCommitTemplate,
		unsetCommitTemplate:   _unsetCommitTemplate,
	}
}

// GetCommitTemplatePath returns the current commit.template configuration
// and whether it provides from global configuration or not.
func (cli *CLI) GetCommitTemplatePath() (path string, global bool, err error) {
	var output bytes.Buffer
	output, err = cli.getCommitTemplatePath()
	if err != nil {
		return
	}

	global = isGlobal(output.Bytes())
	path, err = commitTemplatePath(output.String())
	return
}

// SetCommitTemplate configures Git locally to use a commit template.
func (cli *CLI) SetCommitTemplate(path string) (err error) {
	return cli.setCommitTemplate(path)
}

// UnsetCommitTemplate removes local Git commit template configuration.
func (cli *CLI) UnsetCommitTemplate() (err error) {
	return cli.unsetCommitTemplate()
}

func commitTemplatePath(output string) (string, error) {
	re := regexp.MustCompile(commitTemplatePathRegexp)
	m := re.FindAllStringSubmatch(output, 1)
	if m == nil {
		return "", &NoCommitTemplateConfigurationError{}
	}
	return m[0][1], nil
}

func isGlobal(output []byte) bool {
	re := regexp.MustCompile(globalConfigOutputRegexp)
	return re.Match(output)
}

func _getCommitTemplatePath() (out bytes.Buffer, err error) {
	cmd := exec.Command("git", "config", "--show-origin", "--get", "commit.template")

	cmd.Stdout = &out
	err = cmd.Run()
	return
}

func _setCommitTemplate(path string) (err error) {
	err = _unsetCommitTemplate()
	if err != nil {
		return
	}
	cmd := exec.Command("git", "config", "--add", "commit.template", path)
	err = cmd.Run()
	if err != nil {
		return
	}
	return
}

func _unsetCommitTemplate() (err error) {
	cmd := exec.Command("git", "config", "--unset", "commit.template")
	_ = cmd.Run()
	return
}
