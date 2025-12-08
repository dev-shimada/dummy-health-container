# dummy-health-container

シンプルなヘルスチェックエンドポイントを提供するGo製のWebサーバーです。

## 概要

このアプリケーションは、Echo フレームワークを使用してヘルスチェックエンドポイントを提供する軽量なHTTPサーバーです。開発環境にはDevContainerを使用しています。

## 機能

- カスタマイズ可能なヘルスチェックエンドポイント
- 環境変数によるポート番号の設定
- 構造化ログ出力

## 必要要件

- Go 1.23以上
- Docker（DevContainer使用時）

## 環境変数

| 変数名        | デフォルト値 | 説明                               |
| ------------- | ------------ | ---------------------------------- |
| `HEALTH_PATH` | `/health`    | ヘルスチェックエンドポイントのパス |
| `PORT`        | `8080`       | サーバーが起動するポート番号       |

## 使用方法

### ローカル実行

```bash
go run main.go
```

### DevContainer使用

1. VS Codeでプロジェクトフォルダーを開く
2. コマンドパレット（Cmd/Ctrl + Shift + P）から「Dev Containers: Reopen in Container」を選択
3. コンテナ内でサーバーを起動

```bash
go run main.go
```

### 環境変数を指定して実行

```bash
HEALTH_PATH=/status PORT=3000 go run main.go
```

## エンドポイント

### ヘルスチェック

```
GET /health (デフォルト)
```

**レスポンス:**
```
200 OK
```

## 依存関係

- [Echo](https://echo.labstack.com/) - 高速で最小限のGoウェブフレームワーク

## ライセンス

このプロジェクトはオープンソースソフトウェアです。
