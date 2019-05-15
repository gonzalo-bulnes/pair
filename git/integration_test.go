package git

import (
	"bytes"
	"errors"
	"testing"
)

const exampleGlobalConfigOutput = "file:/home/alice/.gitconfig	/home/alice/.pair"
const exampleLocalConfigOutput = "file:.git/config	/home/bob/.pair"

func TestIntegration(t *testing.T) {

	t.Run("GetCommitTemplatePath()", func(t *testing.T) {

		t.Run("when Git command returns an error", func(t *testing.T) {
			cli := NewCLI()
			var emptyBuffer bytes.Buffer
			cli.getCommitTemplatePath = func() (bytes.Buffer, error) {
				return emptyBuffer, errors.New("Git command failed")
			}

			_, _, err := cli.GetCommitTemplatePath()
			if err == nil {
				t.Error("Expected error, got none")
			}
		})

		t.Run("when Git command succeeds and returns global config", func(t *testing.T) {
			cli := NewCLI()
			output := bytes.NewBufferString(exampleGlobalConfigOutput)
			cli.getCommitTemplatePath = func() (bytes.Buffer, error) {
				return *output, nil
			}

			path, global, err := cli.GetCommitTemplatePath()
			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}
			if path != "/home/alice/.pair" {
				t.Errorf("Unexpected error %v", path)
			}
			if !global {
				t.Error("Expected global flag to be true, got false")
			}
		})

		t.Run("when Git command succeeds and returns repository config", func(t *testing.T) {
			cli := NewCLI()
			output := bytes.NewBufferString(exampleLocalConfigOutput)
			cli.getCommitTemplatePath = func() (bytes.Buffer, error) {
				return *output, nil
			}

			path, global, err := cli.GetCommitTemplatePath()
			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}
			if path != "/home/bob/.pair" {
				t.Errorf("Unexpected error %v", path)
			}
			if global {
				t.Error("Expected global flag to be false, got true")
			}
		})
	})

	t.Run("SetCommitTemplate()", func(t *testing.T) {

		t.Run("when Git command returns an error", func(t *testing.T) {
			cli := NewCLI()
			cli.setCommitTemplate = func(path string) error {
				return errors.New("Git command failed")
			}

			err := cli.SetCommitTemplate("path")
			if err == nil {
				t.Error("Expected error, got none")
			}
		})
	})

	t.Run("UnsetCommitTemplate()", func(t *testing.T) {

		t.Run("when Git command returns an error", func(t *testing.T) {
			cli := NewCLI()
			cli.unsetCommitTemplate = func() error {
				return errors.New("Git command failed")
			}

			err := cli.UnsetCommitTemplate()
			if err == nil {
				t.Error("Expected error, got none")
			}
		})
	})
}
