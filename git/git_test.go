package git_test

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/gonzalo-bulnes/pair/git"
	"github.com/gonzalo-bulnes/pair/template"
)

func ExampleCommitTemplate_CoAuthor_present() {
	tt := git.NewCommitTemplate()

	template := `Add secret message

  Co-Authored-By: Alice <alice@example.com>`
	var buffer bytes.Buffer
	buffer.WriteString(template)
	tt.ReadFrom(&buffer)

	coAuthor, present := tt.CoAuthor()
	fmt.Println(coAuthor)
	fmt.Println(present)
	// Output:
	// Alice <alice@example.com>
	// true
}

func ExampleCommitTemplate_CoAuthor_absent() {
	tt := git.NewCommitTemplate()

	template := `Add secret message`
	var buffer bytes.Buffer
	buffer.WriteString(template)
	tt.ReadFrom(&buffer)

	coAuthor, present := tt.CoAuthor()
	fmt.Println(coAuthor)
	fmt.Println(present)
	// Output:
	// false
}

func TestCommitTemplate(t *testing.T) {
	t.Run("implements T", func(t *testing.T) {
		var _ template.T = (*git.CommitTemplate)(nil)
	})

	t.Run("AddCoAuthor", func(t *testing.T) {
		testcases := []struct {
			templatePath string
			coAuthor     string
			resultPath   string
		}{
			{
				"none.txt",
				"Alice <alice@example.com>",
				"simple.txt",
			},
			{
				"simple.txt",
				"Bob <bob@example.com>",
				"double.txt",
			},
			// unsupported edge case: co-author lines that are not at the end of the template
		}

		for _, tc := range testcases {
			tt := git.NewCommitTemplate()

			f, err := os.Open(filepath.Join("testdata", tc.templatePath))
			if err != nil {
				t.Fatalf("Missing test data: %s", tc.templatePath)
			}
			tt.ReadFrom(f)

			r, err := os.Open(filepath.Join("testdata", tc.resultPath))
			if err != nil {
				t.Fatalf("Missing test data: %s", tc.resultPath)
			}
			var result bytes.Buffer
			result.ReadFrom(r)

			ok := tt.AddCoAuthor(tc.coAuthor)
			if !ok {
				t.Error("Failed to add co-author")
			}
			if tt.String() != result.String() {
				t.Errorf("Expected resulting template to be '%s', was '%s'", result.String(), tt.String())
			}
		}
	})

	t.Run("CoAuthor", func(t *testing.T) {
		testcases := []struct {
			templatePath string
			coAuthor     string
			present      bool
		}{
			{
				"simple.txt",
				"Alice <alice@example.com>",
				true,
			},
			{
				"none.txt",
				"",
				false,
			},
			{
				"double.txt", // only one co-author is supported
				"Alice <alice@example.com>",
				true,
			},
		}

		for _, tc := range testcases {
			tt := git.NewCommitTemplate()

			f, err := os.Open(filepath.Join("testdata", tc.templatePath))
			if err != nil {
				t.Fatalf("Missing test data: %s", tc.templatePath)
			}
			tt.ReadFrom(f)

			coAuthor, present := tt.CoAuthor()
			if present != tc.present {
				t.Errorf("Co-author detection failed for %s", tc.templatePath)
			}
			if coAuthor != tc.coAuthor {
				t.Errorf("Expected co-author to be '%s', was '%s'", tc.coAuthor, coAuthor)
			}
		}
	})
}
