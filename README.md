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

### Priority: ☆☆☆

- どの version までをサポートするか決める
  - 全部をサポートし続けるのは無理。最新3つくらい？
    - バージョン移行用の変換スクリプトも作る？
  - バージョンごとに処理が違う。バージョンによって処理を分岐させるので、バージョンごとにディレクトリを分けてコード書く必要がありそう
  - 面倒すぎる。かなり複雑になる。まともに管理できないし、したくない
  - やはり、バージョンごとに branch を切った方が良さそう
    - つまり、原則として最新1つのみサポート。
      - その代わり、バージョン移行用の変換スクリプトを手厚く？
- Implement a skeleton
  - 処理の大まかな流れ：
    - プロジェクトのルートディレクトリ(？)にある `.fmconfig.toml` を読み込む
      - `.fmconfig.toml` に書かれていること：
      - `path-to-allowed-taxonomies: "foo"`: `allowed-taxonomies.toml` へのパス
      - `path-to-targets: ["foo"]`: 対象となる Markdown ファイルたちのリスト。ディレクトリ名(または直接ファイル名)で指定する。
    - ディレクトリパスを巡回して、ファイルリストを取得
    - 各ファイルの front matter を抽出
    - Golang の構造体へと変換
    - ソートする、一覧表示する、など様々
- `v0.1.0` の仕様を固める
  - `template.md`(?) とその仕様をまとめた `specification.md` を作る
- Update TODO
- Upgrade Golang version from 1.16 to 1.19

### Priority: ☆☆

### Priority: ☆

- GUI(ブラウザ)版作る？(TypeScript+React.jsあたりで書いて、GitHub Pages上で動くやつ)
  - TypeScript, React.js の勉強になってよさそう
  - front matter の仕様が固まり、CLI が完成してからなので、かなり後になりそう。

### Uncateg

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
