# go-othello

## 概要

ターミナルで遊ぶオセロ。

## 仕様

- 8 \* 8 の盤面が出てくる
- ゲーム開始時に白 2 黒 2 が置かれている
- 盤面で自分と違う色の石の隣に黒 or 白の石を置ける
- 違う色で挟まれたらところに石を置かれたら、挟まれた石は色が反転する
- 終了時に勝敗を判定

## 必要
Go

### エディタ
VS Code

### 拡張機能
- [Go](https://marketplace.visualstudio.com/items?itemName=golang.Go)

### ツール(go install)
- [gofumpt](https://github.com/mvdan/gofumpt)
- [gosec](https://github.com/securego/gosec)
- [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports)
- [gopls](https://github.com/golang/tools/blob/master/gopls/README.md)

## 環境準備
```
go mod tidy
```

## テスト方法

## ビルド方法

## デバッグ方法

## 実行方法
```
go run othello.go
```
