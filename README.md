# 前提
以下のコマンドで環境構築を行うことが可能です。ただし、これらのコマンドはLinuxやMacなどのUnix系OSで動作することを前提としています。Windows環境では、適切な変更が必要になるかもしれません。

# バックエンドディレクトリの準備
go-templateディレクトリでgoの初期化を行います。
```
~/home $ cd go-template
~/home/go-template $ go mod init go-template 
~/home/go-template $ go: creating new go.mod: module go-template
~/home/go-template $ ls
go.mod
```

次に基本的なディレクトリとファイルを作成します。
```
~/home/go-template $ mkdir -p home/{controller,database,helper,middleware,model,service} env logs
~/home/go-template $ touch README.md main.go env/env.go env/local.env
```

Ginをインストールします。
```
~/home/go-template $ go get -u github.com/gin-gonic/gin
~/home/go-template $ go get -u github.com/gin-gonic/gin
go: downloading golang.org/x/net v0.10.0
go: downloading golang.org/x/crypto v0.9.0
go: added github.com/bytedance/sonic v1.8.8
go: added github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311
go: added github.com/gin-contrib/sse v0.1.0
go: added github.com/gin-gonic/gin v1.9.0
go: added github.com/go-playground/locales v0.14.1
go: added github.com/go-playground/universal-translator v0.18.1
go: added github.com/go-playground/validator/v10 v10.13.0
go: added github.com/goccy/go-json v0.10.2
go: added github.com/json-iterator/go v1.1.12
go: added github.com/klauspost/cpuid/v2 v2.2.4
go: added github.com/leodido/go-urn v1.2.4
go: added github.com/mattn/go-isatty v0.0.18
go: added github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
go: added github.com/modern-go/reflect2 v1.0.2
go: added github.com/pelletier/go-toml/v2 v2.0.7
go: added github.com/twitchyliquid64/golang-asm v0.15.1
go: added github.com/ugorji/go/codec v1.2.11
go: added golang.org/x/arch v0.3.0
go: added golang.org/x/crypto v0.9.0
go: added golang.org/x/net v0.10.0
go: added golang.org/x/sys v0.8.0
go: added golang.org/x/text v0.9.0
go: added google.golang.org/protobuf v1.30.0
go: added gopkg.in/yaml.v3 v3.0.1
~/home/go-template $ 
```

これでgo.modに必要なライブラリなどがインストールされます。

```go:go.mod
module go-template

go 1.20

require (
	github.com/bytedance/sonic v1.8.8 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.9.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.13.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.7 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

```

これで以下のようになります。
```
.
├── README.md  - プロジェクトの概要、使用技術、セットアップ方法などを記述するためのマークダウンファイルです。
│
├── home  - アプリケーションのコアとなる部分を格納するディレクトリです。
│   ├── controller  - ユーザーのリクエストを受け取り、レスポンスを返す役割を果たすコントローラーを定義するディレクトリです。
│   ├── database  - データベースへのアクセスや操作を管理するためのファイルを格納するディレクトリです。
│   ├── helper  - ヘルパーメソッドやユーティリティ関数を定義するディレクトリです。再利用可能な共通のロジックを格納します。
│   ├── middleware  - リクエストとレスポンスの間に挟んで特定の処理を行うミドルウェアを定義するディレクトリです。
│   ├── model  - データベースのテーブルを表すモデルを定義するディレクトリです。
│   └── service  - ビジネスロジックを実装するサービス層を定義するディレクトリです。
│
├── env  - 環境変数の管理に関するディレクトリです。
│   ├── env.go  - 環境変数を扱うGoのファイルです。ここで環境変数を読み込んだり、デフォルト値を設定します。
│   └── local.env  - ローカル環境の環境変数を設定するためのファイルです。ここにはデータベースの接続情報やAPIキーなどを設定します。
│
├── go.mod  - Goのモジュールとその依存関係を定義するファイルです。
│
└── main.go  - アプリケーションのエントリーポイントとなるGoのファイルです。アプリケーションの初期化と実行を担当します。
```


main.goで以下のように記載する。

```go:main.go
package main

import "github.com/gin-gonic/gin"

import "net/http"

func main() {
    engine:= gin.Default()
    engine.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "hello world",
        })
    })
    engine.Run(":8080")
}

```

以下のコマンドでgoサーバーを起動します。

```
~/home/go-template $ go run main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.main.func1 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
```

[GIN-debug] Listening and serving HTTP on :8080 となっていれば成功

以下のコマンドをターミナル上で実行してレスポンスを確認します。
```
~/home $ curl http://localhost:8080/
{"message":"hello world"}% 
```
hello worldが返ってきたらOK

