// Package github provides primitives for reading and writing co-author declarations.
// See https://help.github.com/articles/creating-a-commit-with-multiple-authors
package github

// CoAuthorRegexp matches Github's co-author declarations and
// captures the co-authors' names.
const CoAuthorRegexp = `\n[Cc]o-[Aa]uthored-[Bb]y: (.*)`

// CoAuthorPrefix starts a co-author declaration.
const CoAuthorPrefix = "\nCo-Authored-By: "
