# いろいろメモ

## hello.goプログラム実行
`go run hello.go`

## 実行ファイル作成と実行
hello.goをhelloとして実行ファイルを作成

`go build -o hello hello.go`

実行ファイルからプログラム実行

`./hello`

ちなみに`go build`でビルドできる。この場合`go build *.go`の動作をする

## パッケージについて
go langでは「1つのディレクトリに1つのパッケージ定義のみ」という原則がある。
そのため複数のパッケージ定義を同一ディレクトリ内におくことは原則としてできないと考えましょう。

## テストについて
テストコードは読んでね(animals_test.go)
テスト実行は`go test ./animals`
テストの詳細を見たい場合は`go test -v ./animals`