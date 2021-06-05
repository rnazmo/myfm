# myfm

myfm (My FrontMatter) is a Golang package to manage front matters for my own use.

## TODO

- [ ] Add a new script `/devel-tools/script/install-devel-tools.linux.x64.sh`
- [ ] Use new GitHub Actions workflows
  - [ ] integ-test
  - [ ] lint integ-test
- [ ] Add `_example/*`
- [ ] Add `testdata/a.md`, `testdata/b.md`, ...
- [ ] Add CLI. (under `/cli/myfm/`, using cobra, commands: extract, format, lint)
- [ ] Write documents
- [ ] Add functions
  - [ ] `Format`
  - [ ] `Validate` (including `validateFrontMatterVersion`, `validateTitle`, ...) (Check as a 'FrontMatter')
- [ ] Add struct `FrontMatter`
- [ ] Add tests
