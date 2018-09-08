// Package pair provides primitives to manage co-author declarations in Git commit templates.
package pair

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

const version = "0.1.0" // adheres to semantic versioning

// Version prints the package version.
func Version(out, errors io.Writer) error {
	fmt.Fprintf(out, fmt.Sprintf("pair version %s\n", version))
	return nil
}

// With configures Git to use a commit template and adds a co-author declaration
// to that commit template.
func With(out, errors io.Writer, pair string) error {
	template := filepath.Join(os.Getenv("HOME"), ".pair")

	err := configureGit(template)
	if err != nil {
		return err
	}
	overwrite(template, fmt.Sprintf("\n\nCo-Authored-By: %s\n", pair))
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

func overwrite(template, pair string) (err error) {
	_ = remove(template)
	f, err := os.OpenFile(template, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	if _, err = f.Write([]byte(pair)); err != nil {
		return
	}
	if err = f.Close(); err != nil {
		return
	}
	return
}

func remove(template string) (err error) {
	err = os.Remove(template)
	if err != nil {
		return err
	}
	return
}

func unconfigureGit() error {
	cmd := exec.Command("git", "config", "--global", "--unset", "commit.template")
	_ = cmd.Run()
	return nil
}
