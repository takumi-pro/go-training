## goコマンド

### go run
go runコマンドは以下のような処理が行われている
- go run実行
- バイナリファイルが一時ディレクトリに格納される
- そのバイナリファイルを実行する
- 実行後にバイナリファイルが削除される

helloworldを出力するプログラムを書いて`go run`を実行

```go
package main

import "training/chapter01"

func main() {
	chapter01.Main()
}
```

### go build
バイナリファイルを作成するときに`go build`を使用する

上記のプログラムをビルドすると同階層にファイル名でバイナルファイルが作成される
ファイル名が`main.go`だとすると`main`ファイルが出力される

### go install
ライブラリのインストールに使用される。
ライブラリは`$GOPATH/bin`配下に格納されることになる。
Go1.15までは`go get`でライブラリのインストールを行なっていたが、`go.mod`を編集する役割も担っていたためgo1.16からinstallとgetでそれぞれ役割が分断された。
