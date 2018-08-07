How does it work?
=================

Github supports [creating commits with multiple authors][gh-docs]. Git supports using [template for commit messages][git-docs].

Pair creates a template for your commit messages that contains the line: `Co-Authored-By: Your Pair <awesome@example.com>`. And it configures Git to use that template.

When you commit, the line is written for you. You can edit, or remove it, and you can change the default name and email by swapping pairs!

It's nothing more than that, a file and a tiny program that re-writes it from time to time.

  [gh-docs]: https://help.github.com/articles/creating-a-commit-with-multiple-authors/
  [git-docs]: https://robots.thoughtbot.com/better-commit-messages-with-a-gitmessage-template

<br /><br />

<p align='center'><img width="128" src='../vendor/noto-emoji-pear.png' alt="A pear emoji." /></p>
