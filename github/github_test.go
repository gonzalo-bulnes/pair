package github_test

import (
	"regexp"
	"testing"

	"github.com/gonzalo-bulnes/pair/github"
)

func TestCoAuthorRegexp(t *testing.T) {
	testcases := []struct {
		source  string
		matches [][]string
	}{
		{
			source: "Some message\n\nCo-Authored-By: Bob <bob@example.com>\n",
			matches: [][]string{
				{
					"Co-Authored-By: Bob <bob@example.com>",
					"Bob <bob@example.com>",
				},
			},
		},
		{
			source:  "Some message\n",
			matches: [][]string{},
		},
		{
			source: "Some message\n\nCo-authored-By: Bob <bob@example.com>\nco-Authored-by: Alice <alice@example.com>",
			matches: [][]string{
				{
					"Co-authored-By: Bob <bob@example.com>",
					"Bob <bob@example.com>",
				},
				{
					"co-Authored-by: Alice <alice@example.com>",
					"Alice <alice@example.com>",
				},
			},
		},
	}

	for _, tc := range testcases {
		re := regexp.MustCompile(github.CoAuthorRegexp)
		matches := re.FindAllStringSubmatch(tc.source, -1)

		if len(matches) != len(tc.matches) {
			t.Fatalf("Expected %d matches, got %d", len(tc.matches), len(matches))
		}
		for i, match := range matches {
			if match[0] != tc.matches[i][0] || match[1] != tc.matches[i][1] {
				t.Errorf("Unexpected match, expected %v, got %v", tc.matches[i], matches)
			}
		}
	}
}
