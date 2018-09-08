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

// GetCommitTemplatePath returns the current commit.template configuration
// and whether it provides from global configuration or not.
func GetCommitTemplatePath() (path string, global bool, err error) {
	cmd := exec.Command("git", "config", "--show-origin", "--get", "commit.template")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()

	global = isGlobal(out.Bytes())
	path, err = commitTemplatePath(out.String())
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
