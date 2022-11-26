# othello-go

## 概要
ターミナルで遊ぶオセロ。

## 仕様
- 8 \* 8 の盤面が出てくる
- ゲーム開始時に白 2 黒 2 が置かれている
- 盤面で自分と違う色の石の隣に黒 or 白の石を置ける
- 違う色で挟まれたらところに石を置かれたら、挟まれた石は色が反転する
- 終了時に勝敗を判定

## 必要
Go(1.19以上)

### エディタ
VS Code

### 拡張機能
- [Go](https://marketplace.visualstudio.com/items?itemName=golang.Go)

### ツール
- [golangci-lint](https://golangci-lint.run/)

## 環境準備
```
go mod tidy
```

## 注意
Macで作成したため、Windows環境では異なる点があります。
具体的には、makeで実行出来ないため、Makefileで定義されているコマンドをそのまま実行して下さい。

## テスト方法
```
make test
```

## ビルド方法
```
make build
```

## デバッグ方法
```
cmd/othello-go/main.goをアクティブなタブにした状態でVS Codeのデバッグを実行する。
```

## 実行方法
```
make run
```
