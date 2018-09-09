How does it work?
=================

Github supports [creating commits with multiple authors][gh-docs]. Git supports using [templates for commit messages][git-docs].

If you don't already use a template for your commit messages, `pair` creates one for you, that contains the line: `Co-Authored-By: Your Pair <awesome@example.com>`. And it configures Git to use that template.

When you commit, the line is written for you. You can edit, or remove it, and you can change the default name and email by swapping pairs!

It's nothing more than that, a file and a tiny program that amends it from time to time.

----

If you do already use a commit template for your messages, `pair` appends a co-author declaration in the form of `Co-Authored-By: Your Pair <awesome@example.com>` to it, and removes it when you stop pairing.

_Note: Things go wild if your commit template already contains a co-author declaration that is not the last line in that template. Should that be your case, please [open an issue][issue] and we can pair in making `pair` support your use case!_


  [gh-docs]: https://help.github.com/articles/creating-a-commit-with-multiple-authors/
  [git-docs]: https://robots.thoughtbot.com/better-commit-messages-with-a-gitmessage-template
  [issue]: https://github.com/gonzalo-bulnes/pair/issues

<br /><br />

<p align='center'><img width="128" src='../vendor/noto-emoji-pear.png' alt="A pear emoji." /></p>
