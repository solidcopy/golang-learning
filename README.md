# このリポジトリについて

## 概要

Go言語について学習したことを動作するコードとコメントの形で残す。

## 構成

テーマごとに細かくディレクトリと `main.go` を作る。
どのディレクトリも `main` パッケージなので個別に実行できる。

コードで表現できないことはこのファイルにまとめる。

## 実行

`hello` を実行する場合

`$ go run ./cmd/hello`

# 目次

- 初めの一歩
  - hello
    - 最小構成
- モジュール
  - imports
    - インポート
  - init
    - パッケージ初期化
  - import_init_only
    - initのみ実行
- データ
  - data_types
    - データ型
  - cast_and_convert
    - キャストと変換
  - variables
    - 変数
  - zero_values
    - ゼロ値
  - constants
    - 定数
  - constants_sequences
    - 列挙定数
  - pointers
    - ポインタ
  - structs
    - 構造体
  - methods
    - メソッド
  - interfaces
    - インタフェース
  - arrays_and_slices
    - 配列とスライス
  - maps
    - マップ
  - new_and_make
    - newとmake
- 制御文
  - conditional
    - 条件分岐
  - loops
    - 繰り返し
- 関数
  - functions
    - 関数
  - function_errors
    - エラーを返す関数
- エラーハンドリング
  - errors
    - error型
  - panics
    - パニック
  - defer
    - deferによる後処理
- 並列処理
  - goroutines
    - Goルーチン
  - mutexes
    - Mutexによる排他制御
- テスト
  - unit_tests
    - 単体テスト
- ドキュメンテーション
  - documentation
    - コメントのドキュメント化
- ログ
  - logs
    - ログ
- 文字列処理
  - regexp
    - 正規表現

# コード化できない知識

## Goのインストール

公式バイナリを展開して `/usr/local/go` に配置する。

PATHを通す。

`$ go version`

で動作確認する。

## Goのアップグレード

`/usr/local/go` の中身を削除してインストールの手順を繰り返す。

プロジェクトのGoバージョン要件を最新に変更するなら

`$ go get go@latest`

マイナーバージョンを変えず最新パッチに更新するなら

`$ go get toolchain@patch`

で `go.mod` を更新する。

## モジュールとパッケージ

モジュールは開発するアプリやライブラリなど、バージョン管理の単位。

パッケージはディレクトリ。
インポートの単位はパッケージ。

パッケージ内では関数、型、変数、定数が共有される。

パッケージ名とディレクトリ名は基本的に一致させるが必須ではない。
このリポジトリでも `main` パッケージを様々なディレクトリ名で作っている。

## モジュールの作成と管理

### モジュールの新規作成

モジュールにするディレクトリ内で、

`go mod init <module-path>`

`<module-path>` はモジュールを一意に特定する。
GitHubリポジトリ、またはダウンロードできるものならそのURLなど。

例）
`go mod init com.github/solidcopy/golang-learning`

モジュールの依存関係などを記録する `go.mod` ができる。

## モジュールのディレクトリ構成

## モジュールの利用

モジュールパスとパッケージのディレクトリパスを連結して `import` に書く。

[https://pkg.go.dev/](https://pkg.go.dev/)で配布されているモジュールを探す。

新たなモジュールを利用するなら `import` に書いてから

`go mod tidy`

## 公開していないモジュールの利用

ディレクトリ `../my-module` にあるモジュールを
`github.com/solidcopy/my-module` としてインポートする場合、

`go mod edit -replace=github.com/solidcopy/my-module=../my-module`

`go.mod` にはこう書かれる。

`replace github.com/solidcopy/my-module => ../my-module`

# ビルド

エントリポイントのパッケージが `cmd/server` の場合、

`go build ./cmd/server`

## Windows用ビルド

`GOOS=windows GOARCH=amd64 go build -o my_program.exe`

## 開発したツールのインストール

`go install`

ビルドしたバイナリを所定のディレクトリに配置して実行できるようにする。

`go list -f '{{.Target}}'`

で配置先を調査する。

`go env -w GOBIN=C:\path\to\your\bin`

で配置先を変更する。

## テスト実行

`go test ./cmd/unti_tests`
