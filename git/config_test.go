package git_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/gonzalo-bulnes/pair/git"
)

func TestConfig(t *testing.T) {
	t.Run("Apply()", func(t *testing.T) {
		templatePath := full("config", "arbitrary.txt")

		config, err := git.NewConfig(templatePath)
		if err != nil {
			t.Fatalf("Missing test data: %s", templatePath)
		}
		config.CommitTemplate.AddCoAuthor("Lewis Caroll <lewis@wonderland.example.io>")
		defer func() {
			config, err = git.NewConfig(templatePath)
			if err != nil {
				t.Fatalf("Missing test data: %s", templatePath)
			}
			config.CommitTemplate.RemoveCoAuthor("Lewis Caroll <lewis@wonderland.example.io>")
			config.Apply()
		}()
		expectedContent := config.CommitTemplate.String()
		config.Apply()

		template, err := os.Open(templatePath)
		if err != nil {
			t.Fatalf("Missing test data: %s", templatePath)
		}
		var templateContent bytes.Buffer
		templateContent.ReadFrom(template)

		if templateContent.String() != expectedContent {
			t.Errorf("Expected template file to contain '%s', instead contains '%s'",
				expectedContent, templateContent.String())
		}
	})

	t.Run("NewConfig()", func(t *testing.T) {
		t.Run("with wrong path", func(t *testing.T) {
			_, err := git.NewConfig("missing.txt")
			if err == nil {
				t.Fatalf("Expected error, got none")
			}
		})

		t.Run("with correct path", func(t *testing.T) {
			originalPath := full("config", "arbitrary.txt")

			original, err := os.Open(originalPath)
			if err != nil {
				t.Fatalf("Missing test data: %s", originalPath)
			}
			var originalContent bytes.Buffer
			originalContent.ReadFrom(original)

			config, err := git.NewConfig(originalPath)
			if err != nil {
				t.Fatalf("Missing test data: %s", originalPath)
			}

			if config.CommitTemplatePath != originalPath {
				t.Errorf("Expected commit template path to be %s, was %s", originalPath, config.CommitTemplatePath)
			}

			if config.CommitTemplate.String() != originalContent.String() {
				t.Errorf("Expected commit template path to be %s, was %s", originalContent.String(), config.CommitTemplate.String())
			}
		})
	})
}
