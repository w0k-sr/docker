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
docker build -t market_app .

# 実行（CSVファイルに出力）
# 仮想通貨名:bitcoin,ripple
# 取得開始日、取得終了日：yyyy-mm-dd
docker run --rm -v $(pwd)/output:/app/output market_app -coin=<仮想通貨名> -start=<取得開始日> -end=<取得終了日>