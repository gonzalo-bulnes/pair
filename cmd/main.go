// Command pair allows to set and manage co-author declarations
// for seamless pairing sessions using Git and Github.
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/gonzalo-bulnes/pair"
)

type config struct {
	templateName string
	pair         string
}

type argumentsError struct{}

func (e *argumentsError) Error() string {
	return `Usage:

pair with <email>
Example:
  pair with 'Alice <alice@example.com>'

pair stop
`
}

func main() {
	err := checkArgs(os.Args)
	if err != nil {
		os.Exit(fail(os.Stderr, err, 0))
	}

	git := pair.GetGitConnector()

	switch os.Args[1] {
	case "--version":
		_ = pair.Version(os.Stdout, os.Stderr)
		os.Exit(0)
	case "with":
		err = pair.With(git, os.Stdout, os.Stderr, os.Args[2])
		if err != nil {
			os.Exit(fail(os.Stderr, err, 20))
		}
		os.Exit(0)
	case "stop":
		err = pair.Stop(git, os.Stdout, os.Stderr)
		if err != nil {
			os.Exit(fail(os.Stderr, err, 21))
		}
		os.Exit(0)
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
	if len(args) == 2 && args[1] == "--version" {
		return nil
	}
	return &argumentsError{}
}

func fail(errors io.Writer, err error, code int) int {
	if e, ok := err.(*argumentsError); ok {
		fmt.Fprint(errors, e)
		return code
	}
	var version bytes.Buffer
	pair.Version(&version, errors)

	fmt.Fprintf(errors, `
Oh no! You might have found a bug in pair!

Please open an issue mentioning (error %d) and maybe we can pair to fix it : )
https://github.com/gonzalo-bulnes/pair/issues

%serror: %v
`, code, version.String(), err)
	return code
}
