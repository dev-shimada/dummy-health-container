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

### Docker コンテナ（GHCR）

GHCR (GitHub Container Registry) からイメージを pull して実行できます。

```bash
docker pull ghcr.io/dev-shimada/dummy-health-container:latest
```

#### 基本的な実行

```bash
docker run -p 8080:8080 ghcr.io/dev-shimada/dummy-health-container:latest
```

#### 環境変数を指定して実行

```bash
docker run -p 3000:3000 \
  -e HEALTH_PATH=/status \
  -e PORT=3000 \
  ghcr.io/dev-shimada/dummy-health-container:latest
```

#### rootful

```bash
docker pull ghcr.io/dev-shimada/dummy-health-container:latest-rootful
```

```bash
docker run -p 80:80 \
  -e PORT=80 \
  ghcr.io/dev-shimada/dummy-health-container:latest-rootful
```

#### タグ

| タグ                                  | 説明                               |
| ------------------------------------- | ---------------------------------- |
| `latest`                              | 最新の安定版（non-root）           |
| `latest-rootful`                      | 最新の安定版（root）               |
| `v1.0.0` など                         | 特定バージョン（non-root）         |
| `v1.0.0-rootful`                      | 特定バージョン（root）             |
| `v1.0.0-linux-amd64`                  | amd64 アーキテクチャ向け           |
| `v1.0.0-linux-arm64`                  | arm64 アーキテクチャ向け           |
| `v1.0.0-rootful-linux-amd64`          | amd64 アーキテクチャ向け（root）   |
| `v1.0.0-rootful-linux-arm64`          | arm64 アーキテクチャ向け（root）   |

### ローカルで Docker ビルド

```bash
docker build -t dummy-health-container .
```

#### 実行

```bash
docker run -p 8080:8080 dummy-health-container
```

#### 環境変数を指定して実行

```bash
docker run -p 3000:3000 \
  -e HEALTH_PATH=/status \
  -e PORT=3000 \
  dummy-health-container
```

#### rootful

```bash
docker build -f Dockerfile.rootful -t dummy-health-container:rootful .
```

```bash
docker run -p 80:80 \
  -e PORT=80 \
  dummy-health-container:rootful
```

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
