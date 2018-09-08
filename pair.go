// Package pair provides primitives to manage co-author declarations in Git commit templates.
package pair

import (
	"fmt"
	"io"
)

const version = "0.1.0" // adheres to semantic versioning

// Version prints the package version.
func Version(out, errors io.Writer) error {
	fmt.Fprintf(out, fmt.Sprintf("pair version %s\n", version))
	return nil
}
