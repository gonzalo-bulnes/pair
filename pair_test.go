package pair_test

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/gonzalo-bulnes/pair"
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
}
