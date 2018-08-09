package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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
	template := fmt.Sprintf("%s/.pair", home)

	err := checkArgs(os.Args)
	if e, ok := err.(*argumentsError); ok {
		fmt.Println(e)
		return
	}

	switch os.Args[1] {
	case "with":
		err := configureGit(template)
		if err != nil {
			log.Fatal(err)
		}
		pair := os.Args[2]
		overwrite(template, fmt.Sprintf("\n\nCo-Authored-By: %s\n", pair))
		if err != nil {
			log.Fatal(err)
		}
	case "stop":
		_ = remove(template)
		err := unconfigureGit()
		if err != nil {
			log.Fatal(err)
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
	return &argumentsError{}
}

func configureGit(template string) error {
	err := unconfigureGit()
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("git", "config", "--global", "--add", "commit.template", template)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func overwrite(template, pair string) error {
	_ = remove(template)
	f, err := os.OpenFile(template, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(pair)); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func remove(template string) error {
	err := os.Remove(template)
	if err != nil {
		return err
	}
	return nil
}

func unconfigureGit() error {
	cmd := exec.Command("git", "config", "--global", "--unset", "commit.template")
	_ = cmd.Run()
	return nil
}
