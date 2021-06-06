# myfm

[![Test](https://github.com/rnazmo/myfm/actions/workflows/test.yml/badge.svg)](https://github.com/rnazmo/myfm/actions/workflows/test.yml)
[![Lint](https://github.com/rnazmo/myfm/actions/workflows/lint.yml/badge.svg)](https://github.com/rnazmo/myfm/actions/workflows/lint.yml)

myfm (My FrontMatter) is a Golang package to manage front matters for my own use.

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
  - [ ] `func (fm frontmatter) Marshal() (???, error)`: Convert the struct to toml
  - [ ] `func (fm frontmatter) MarshalToJson() (???, error)`: Convert the struct to json
- [ ] (Add struct `FrontMatter`)?
- [ ] Add GitHub Actions workflows
  - [ ] integ-test
- [ ] Add `_example/*`
- [ ] Add `testdata/a.md`, `testdata/b.md`, ...
- [ ] Add CLI. (under `/cli/myfm/`, using cobra, commands: extract, format, lint)
- [ ] Write documents
- [ ] Make `/internal/` directory and move `formatter` package to under it.
- [ ] Make the field names of `frontmatter` non-capitalizable.
  - It prevents the field values from being changed from outside the package. Then the field values of `frontmatter` are always guaranteed to be validated.
  - We access the field values via methods like `frontMatterVersion() (string)`, `SetFrontMatterVersion(frontMatterVersion string) error`.
  - But.... We cannot unmarshal the `frontmatter` struct if the field names are not capitalized.
