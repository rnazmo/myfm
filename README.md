# myfm

[![Test](https://github.com/rnazmo/myfm/actions/workflows/test.yml/badge.svg)](https://github.com/rnazmo/myfm/actions/workflows/test.yml)
[![Lint](https://github.com/rnazmo/myfm/actions/workflows/lint.yml/badge.svg)](https://github.com/rnazmo/myfm/actions/workflows/lint.yml)

myfm (My FrontMatter) is a Golang package to manage front matters for my own use.

## TODO

- [ ] Implement `formatter` package
  - [ ] Implement functions
  - [ ] Add tests
- [ ] Add scripts
  - [ ] `/devel-tools/script/install-devel-tools.linux.x64.sh`
  - [ ] Add test cases
- [ ] Use Dependabot
- [ ] Add functions
  - [ ] `New(post []byte) (fm FrontMatter, content []byte, err error)`
  - [ ] `Validate` (including `validateFrontMatterVersion`, `validateTitle`, ...) (Check as a 'FrontMatter')
  - [ ] `Format`
- [ ] Add struct `FrontMatter`
- [ ] Add GitHub Actions workflows
  - [ ] integ-test
- [ ] Add `_example/*`
- [ ] Add `testdata/a.md`, `testdata/b.md`, ...
- [ ] Add CLI. (under `/cli/myfm/`, using cobra, commands: extract, format, lint)
- [ ] Write documents
