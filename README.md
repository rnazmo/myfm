# myfm

[![Test](https://github.com/rnazmo/myfm/actions/workflows/test.yml/badge.svg)](https://github.com/rnazmo/myfm/actions/workflows/test.yml)
[![Lint](https://github.com/rnazmo/myfm/actions/workflows/lint.yml/badge.svg)](https://github.com/rnazmo/myfm/actions/workflows/lint.yml)

myfm (My FrontMatter) is a Golang package to manage front matters for my own use.

## Documentation for users

TODO:

## Documentation for developers

### How to setup your development environment

TODO:

## TODO

- [ ] Implement `formatter` package
  - [ ] Implement functions
  - [ ] Add tests
- [ ] Add tests for functions
  - [ ] `NewFromPost`
  - [ ] `NewFromInputs`
  - [ ] `validate`
- [ ] Add test cases for functions
  - [ ] `unmarshal`
- [ ] Add scripts
  - [ ] `/devel-tools/script/install-devel-tools.linux.x64.sh`
- [ ] Use Dependabot
- [ ] Add functions
  - [ ] `New(post []byte) (fm FrontMatter, content []byte, err error)`
  - [ ] `Validate` (including `validateFrontMatterVersion`, `validateTitle`, ...) (Check as a 'FrontMatter')
  - [ ] `Format`
- [ ] Add methods to struct `frontmatter`
  - [ ] `func (fm frontmatter) MarshalToJson() ([]byte, error)`: Convert the struct to json
- [ ] Add GitHub Actions workflows
  - [ ] integ-test
- [ ] Add `_example/*`
- [ ] Add `testdata/a.md`, `testdata/b.md`, ...
- [ ] Add CLI. (under `/cli/myfm/`, using cobra, commands: extract, format, lint)
- [ ] Write documents
- Add `type Frontmatter interface` ?
- [ ] Release `v0.1.0` (or `v0.0.4`? (is based on the `front_matter_version`))
- [ ] Update the front matter template and bump the version of `front_matter_version` (to `v0.0.5`? or `v2.0.0`?)
- [ ] (Add struct `type Frontmatter frontmatter`)?
- [ ] Add struct `frontmatters []frontmatter`?
  - [ ] Add methods to the struct `frontmatters`
    - [ ] `func (fms frontmatters) Titles() []string`
    - [ ] `func (fms frontmatters) Drafteds() []string`
    - [ ] `func (fms frontmatters) Createds() []string`
    - [ ] `func (fms frontmatters) LastUpdateds() []string`
    - [ ] `func (fms frontmatters) LastCheckeds() []string`
    - [ ] `func (fms frontmatters) Tagss() [][]string`
    - [ ] `func (fms frontmatters) IDs() []string`
- [ ] Add badges to `README.md`
  - [ ] Ref: Badges example:
    - https://github.com/go-playground/validator
    - https://github.com/uber-go/zap
  - [ ] go report: https://goreportcard.com/
  - [ ] coverage (codecov?)
    - https://about.codecov.io/blog/getting-started-with-code-coverage-for-golang/
  - [ ] pkg.go.dev: https://pkg.go.dev/badge/
- [ ] logging?
  - Do we really need logging?
  - No
- [ ] `versions` ディレクトリを作り、各バージョンについて `template.md` とその仕様をまとめた `specification.txt` を作る
