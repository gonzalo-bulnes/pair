package git_test

import "path/filepath"

func full(collection, path string) string {
	return filepath.Join("testdata", collection, path)
}
