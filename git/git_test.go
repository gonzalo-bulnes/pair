package git_test

import (
	"bytes"
	"os"
	"path/filepath"
)

func full(collection, path string) string {
	return filepath.Join("testdata", collection, path)
}

func backup(path string) (bytes.Buffer, error) {
	var b bytes.Buffer
	f, err := os.Open(path)
	if err != nil {
		return b, err
	}
	b.ReadFrom(f)
	return b, nil
}

func restore(backup bytes.Buffer, path string) error {
	_ = os.Remove(path)
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, err = backup.WriteTo(f)
	if err != nil {
		return err
	}
	return nil
}
