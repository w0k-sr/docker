# 仮想通貨の時価一括出力ツール

Goで開発した、主要な仮想通貨の価格情報を外部APIから取得し、CSV形式で一括出力するCLIツールです。Docker対応により、環境構築なしですぐに動作確認が可能です。

## 📌 機能概要

- 仮想通貨の現在価格・変動率などをAPIから取得
- 取得結果をCSVファイルとして出力
- 並列処理により複数通貨を高速に取得
- Dockerで簡単に実行可能

## 🚀 使用技術

- Go
- Docker
- 外部API（例: CoinGecko等） ※実際に使っているAPI名に変更してください
- CSV出力処理

## 🛠️ セットアップ手順

### Docker での実行

```bash
# リポジトリをクローン
git clone https://github.com/w0k-sr/docker.git
cd docker/market

# Dockerビルド
docker-compose up -d --build

# 実行（CSVファイルに出力）
# 仮想通貨名:bitcoin,ripple
# 取得開始日、取得終了日：yyyy-mm-dd
docker exec -it market_app bash
/go/src/app> go run main.go -coin=<仮想通貨名> -start=<取得開始日> -end=<取得終了日>
```
## 📂 ディレクトリ構成
```bash
market/
├ GO
|  ├ myapp/
|  |  ├ main.go
|  |  ├ controlles/
|  |    ├ CoinGeckoClient.go
|  ├ Dockerfile          
└── README.md
```
## 🧪 出力例

**出力ファイル名：`prices_ripple_F20250525_T20250605.csv`**

```csv
date,usd_price,jpy_price
2025-05-25,2.08,327.53
2025-05-26,2.32,364.53
...
