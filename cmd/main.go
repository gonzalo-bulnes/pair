package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gonzalo-bulnes/pair"
)

type config struct {
	templateName string
	pair         string
}

type argumentsError struct {
	message string
}

func (e *argumentsError) Error() string {
	return "Usage: pair with <email>\n\nExample:\n\n  pair with 'Alice <alice@example.com>'\n"
}

func main() {
	home := os.Getenv("HOME")
	template := filepath.Join(home, ".pair")

	err := checkArgs(os.Args)
	if e, ok := err.(*argumentsError); ok {
		fmt.Fprint(os.Stderr, e)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "--version":
		_ = pair.Version(os.Stdout, os.Stderr)
		os.Exit(0)
	case "with":
		_ = pair.With(os.Stdout, os.Stderr, os.Args[2])
		os.Exit(0)
	case "stop":
		_ = remove(template)
		err := unconfigureGit()
		if err != nil {
			return
		}
	default:
	}
}

func checkArgs(args []string) error {
	if len(args) == 3 && args[1] == "with" {
		return nil
	}
	if len(args) == 2 && args[1] == "stop" {
		return nil
	}
	if len(args) == 2 && args[1] == "--version" {
		return nil
	}
	return &argumentsError{}
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
