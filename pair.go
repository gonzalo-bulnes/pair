// Package pair provides primitives to manage co-author declarations in Git commit templates.
package pair

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gonzalo-bulnes/pair/git"
)

const version = "0.1.0" // adheres to semantic versioning

// Stop removes the co-author declaration from the commit template.
func Stop(out, errors io.Writer) error {
	commitTemplatePath := defaultCommitTemplatePath()
	ensureExists(commitTemplatePath)
	config, err := git.NewConfig(commitTemplatePath)
	if err != nil {
		return err
	}

	if author, present := config.CommitTemplate.CoAuthor(); present {
		ok := config.CommitTemplate.RemoveCoAuthor(author)
		if !ok {
			return fmt.Errorf("Unable to remove co-author '%s'", author)
		}
	}
	err = config.Apply()
	if err != nil {
		return err
	}
	return nil
}

// Version prints the package version.
func Version(out, errors io.Writer) error {
	fmt.Fprintf(out, fmt.Sprintf("pair version %s\n", version))
	return nil
}

// With configures Git to use a commit template and adds a co-author declaration
// to that commit template.
func With(out, errors io.Writer, pair string) error {

	err := Stop(out, errors)
	if err != nil {
		return err
	}

	commitTemplatePath := defaultCommitTemplatePath()
	ensureExists(commitTemplatePath)
	config, err := git.NewConfig(commitTemplatePath)
	if err != nil {
		return err
	}

	err = configureGit(commitTemplatePath)
	if err != nil {
		return err
	}

	config.CommitTemplate.AddCoAuthor(pair)
	err = config.Apply()
	if err != nil {
		return err
	}
	return nil
}

func configureGit(template string) (err error) {
	err = unconfigureGit()
	if err != nil {
		return
	}
	cmd := exec.Command("git", "config", "--global", "--add", "commit.template", template)
	err = cmd.Run()
	if err != nil {
		return
	}
	return
}

func defaultCommitTemplatePath() string {
	return filepath.Join(os.Getenv("HOME"), ".pair")
}

func ensureExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_, err := os.Create(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func unconfigureGit() error {
	cmd := exec.Command("git", "config", "--global", "--unset", "commit.template")
	_ = cmd.Run()
	return nil
}
