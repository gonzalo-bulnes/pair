package github

// CoAuthorRegexp matches Github's co-author declarations and
// captures the co-authors' names.
// See https://help.github.com/articles/creating-a-commit-with-multiple-authors
const CoAuthorRegexp = `[Cc]o-[Aa]uthored-[Bb]y: (.*)`
