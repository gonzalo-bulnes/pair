package pair_test

import (
	"bytes"
	"fmt"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/gonzalo-bulnes/pair"
	"github.com/gonzalo-bulnes/pair/git"
)

// semverRegexp matches the most common semantinc version strings.
// See https://semver.org
const semverRegexp = `\s\d+.\d+.\d+(?:-(?:alpha|beta|rc)\d*)?\s`

func TestPair(t *testing.T) {
	t.Run("Version()", func(t *testing.T) {

		var out, error bytes.Buffer
		err := pair.Version(&out, &error)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !regexp.MustCompile(semverRegexp).Match(out.Bytes()) {
			t.Errorf("Expected (semantic) version to be put out, was not")
		}
	})

	t.Run("Stop()", func(t *testing.T) {

		unexpectedErr := fmt.Errorf("some specific error")

		testcases := []struct {
			description                        string
			git                                *MockGitConnector
			expectedGetCommitTemplatePathCount int
			expectedSetCommitTemplateCount     int
			expectedUnsetCommitTemplateCount   int
			expectedError                      error
		}{
			{
				description: `When commit.template can't be fetched, error should be returned`,
				git: &MockGitConnector{
					GetCommitTemplatePathError: unexpectedErr,
				},
				expectedGetCommitTemplatePathCount: 1,
				expectedError:                      unexpectedErr,
			},
			{
				description: `When no commit.template is set, no errors should be returned`,
				git: &MockGitConnector{
					GetCommitTemplatePathError: &git.NoCommitTemplateConfigurationError{},
				},
				expectedGetCommitTemplatePathCount: 1,
				expectedError:                      nil,
			},
			{
				description: `With empty commit.template, no errors should be returned`,
				git: &MockGitConnector{
					GetCommitTemplatePathError: nil,
				},
				expectedGetCommitTemplatePathCount: 1,
				expectedError:                      nil,
			},
			{
				description: `With existing commit.template, no errors should be returned`,
				git: &MockGitConnector{
					CommitTemplatePath:         full("existing.txt"),
					GetCommitTemplatePathError: nil,
				},
				expectedGetCommitTemplatePathCount: 1,
				expectedError:                      nil,
			},
		}

		for _, tc := range testcases {
			t.Run(tc.description, func(t *testing.T) {
				var out, error bytes.Buffer
				err := pair.Stop(tc.git, &out, &error)
				if err != tc.expectedError {
					t.Errorf("Expected error to be %v, got %v", tc.expectedError, err)
				}
				if tc.git.GetCommitTemplatePathCount != tc.expectedGetCommitTemplatePathCount {
					t.Errorf("Expected GetCommitTemplatePath to be called %d time(s), was %d",
						tc.expectedGetCommitTemplatePathCount, tc.git.GetCommitTemplatePathCount)
				}
				if tc.git.SetCommitTemplateCount != tc.expectedSetCommitTemplateCount {
					t.Errorf("Expected SetCommitTemplate to be called %d time(s), was %d",
						tc.expectedSetCommitTemplateCount, tc.git.SetCommitTemplateCount)
				}
				if tc.git.UnsetCommitTemplateCount != tc.expectedUnsetCommitTemplateCount {
					t.Errorf("Expected UnsetCommitTemplate to be called %d time(s), was %d",
						tc.expectedUnsetCommitTemplateCount, tc.git.UnsetCommitTemplateCount)
				}
			})
		}
	})

	t.Run("With()", func(t *testing.T) {

		unexpectedErr := fmt.Errorf("some specific error")

		testcases := []struct {
			description                        string
			git                                *MockGitConnector
			expectedGetCommitTemplatePathCount int
			expectedSetCommitTemplateCount     int
			expectedUnsetCommitTemplateCount   int
			expectedError                      error
		}{
			{
				description: `When commit.template can't be fetched, error should be returned`,
				git: &MockGitConnector{
					GetCommitTemplatePathError: unexpectedErr,
				},
				expectedGetCommitTemplatePathCount: 1,
				expectedSetCommitTemplateCount:     0,
				expectedUnsetCommitTemplateCount:   0,
				expectedError:                      unexpectedErr,
			},
			{
				description: `When no commit.template is set, no errors should be returned`,
				git: &MockGitConnector{
					GetCommitTemplatePathError: &git.NoCommitTemplateConfigurationError{},
				},
				expectedGetCommitTemplatePathCount: 2,
				expectedSetCommitTemplateCount:     1,
				expectedUnsetCommitTemplateCount:   0,
				expectedError:                      nil,
			},
			{
				description: `When commit.template can't be set, error should be returned`,
				git: &MockGitConnector{
					SetCommitTemplateError: unexpectedErr,
				},
				expectedGetCommitTemplatePathCount: 2,
				expectedSetCommitTemplateCount:     1,
				expectedUnsetCommitTemplateCount:   0,
				expectedError:                      unexpectedErr,
			},
			{
				description: `With no errors when fetching commit.template, no errors should be returned`,
				git: &MockGitConnector{
					GetCommitTemplatePathError: nil,
				},
				expectedGetCommitTemplatePathCount: 2,
				expectedSetCommitTemplateCount:     1,
				expectedUnsetCommitTemplateCount:   0,
				expectedError:                      nil,
			},
			{
				description: `With existing commit.template, setting should not be changed`,
				git: &MockGitConnector{
					CommitTemplatePath:         full("simple.txt"),
					GetCommitTemplatePathError: nil,
				},
				expectedGetCommitTemplatePathCount: 2,
				expectedSetCommitTemplateCount:     0,
				expectedError:                      nil,
			},
		}

		for _, tc := range testcases {
			t.Run(tc.description, func(t *testing.T) {
				var out, error bytes.Buffer
				err := pair.With(tc.git, &out, &error, "Alice <alice@example.com>")
				if err != tc.expectedError {
					t.Errorf("Expected error to be %v, got %v", tc.expectedError, err)
				}
				if tc.git.GetCommitTemplatePathCount != tc.expectedGetCommitTemplatePathCount {
					t.Errorf("Expected GetCommitTemplatePath to be called %d time(s), was %d",
						tc.expectedGetCommitTemplatePathCount, tc.git.GetCommitTemplatePathCount)
				}
				if tc.git.SetCommitTemplateCount != tc.expectedSetCommitTemplateCount {
					t.Errorf("Expected SetCommitTemplate to be called %d time(s), was %d",
						tc.expectedSetCommitTemplateCount, tc.git.SetCommitTemplateCount)
				}
				if tc.git.UnsetCommitTemplateCount != tc.expectedUnsetCommitTemplateCount {
					t.Errorf("Expected UnsetCommitTemplate to be called %d time(s), was %d",
						tc.expectedUnsetCommitTemplateCount, tc.git.UnsetCommitTemplateCount)
				}
			})
		}
	})
}

type MockGitConnector struct {
	CommitTemplatePath         string
	CommitTemplatePathGlobal   bool
	GetCommitTemplatePathError error
	SetCommitTemplateError     error
	UnsetCommitTemplateError   error
	GetCommitTemplatePathCount int
	SetCommitTemplateCount     int
	UnsetCommitTemplateCount   int
}

func (m *MockGitConnector) GetCommitTemplatePath() (path string, global bool, err error) {
	m.GetCommitTemplatePathCount += 1
	return m.CommitTemplatePath, m.CommitTemplatePathGlobal, m.GetCommitTemplatePathError
}
func (m *MockGitConnector) SetCommitTemplate(path string) (err error) {
	m.SetCommitTemplateCount += 1
	return m.SetCommitTemplateError
}
func (m *MockGitConnector) UnsetCommitTemplate() (err error) {
	m.UnsetCommitTemplateCount += 1
	return m.UnsetCommitTemplateError
}

func full(path string) string {
	return filepath.Join("testdata", path)
}
