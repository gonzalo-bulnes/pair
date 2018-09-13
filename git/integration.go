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
type CLI struct{}

// GetCommitTemplatePath returns the current commit.template configuration
// and whether it provides from global configuration or not.
func (cli *CLI) GetCommitTemplatePath() (path string, global bool, err error) {
	cmd := exec.Command("git", "config", "--show-origin", "--get", "commit.template")

	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return
	}

	global = isGlobal(out.Bytes())
	path, err = commitTemplatePath(out.String())
	return
}

// SetCommitTemplate configures Git locally to use a commit template.
func (cli *CLI) SetCommitTemplate(path string) (err error) {
	err = cli.UnsetCommitTemplate()
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

// UnsetCommitTemplate removes local Git commit template configuration.
func (cli *CLI) UnsetCommitTemplate() (err error) {
	cmd := exec.Command("git", "config", "--unset", "commit.template")
	_ = cmd.Run()
	return
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
