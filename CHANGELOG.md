# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/) and
this project versions are for now mere release numbers. They don't have any particular meaning.

This project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [1.0.0-alpha][1.0.0-alpha] - 2018-09-09

### Added

- Support for pre-existing commit templates (i.e. you can use both your usual template and pair with ease!)

### Changed

- Git `commit.template` configuration is not overwritten any more.
- If missing both locally and globally, local `commit.template` configuration is set up.
- Editing you commit template won't disturb your `pair` usage, just do as usual.

## 0.1.0 - 2018-08-16

### Added

- Proof of concept implementation of `pair`

> **CAUTION**: This version does overwrite your global Git commit.template configuration.

  [1.0.0-alpha]: https://github.com/gonzalo-bulnes/pair/compare/v0.1.0...add-support-for-preexisting-commit-template
