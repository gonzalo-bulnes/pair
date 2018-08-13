package main

import (
	"fmt"
	"os"
	"os/exec"
)

const Version = "0.1.0"

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
	template := fmt.Sprintf("%s/.pair", home)

	err := checkArgs(os.Args)
	if e, ok := err.(*argumentsError); ok {
		fmt.Fprint(os.Stderr, e)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "--version":
		fmt.Printf("%s\n", version())
		os.Exit(0)
	case "with":
		err := configureGit(template)
		if err != nil {
			return
		}
		pair := os.Args[2]
		overwrite(template, fmt.Sprintf("\n\nCo-Authored-By: %s\n", pair))
		if err != nil {
			return
		}
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

func version() string {
	return fmt.Sprintf("pair version %s", Version)
}
