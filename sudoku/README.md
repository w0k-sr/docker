# 数独アプリケーション（Go + Beego）

このプロジェクトは、GoのWebフレームワーク「Beego」を用いて開発した**数独アプリケーション**です。Docker環境で動作可能な構成とし、UIにはjQueryとCSSを使用して直感的な操作性を実現しています。データの永続化には PostgreSQL を利用しています。

---

## 🧩 主な機能

- 数独の問題生成と入力UI
- 入力値の解答チェック
- 数独の解答アルゴリズム
- 問題の保存（PostgreSQL）

---

## 🛠️ 技術構成

| 技術         | 説明                          |
|--------------|-------------------------------|
| Go           | サーバーサイド言語             |
| Beego        | Go製Webアプリケーションフレームワーク |
| jQuery       | クライアント操作の制御         |
| CSS          | UIスタイルの整形               |
| PostgreSQL   | データベース                    |
| Docker       | 開発・実行環境のコンテナ化       |
| VSCode       | 開発エディタ                   |

---

## 🚀 起動方法

### 1. リポジトリのクローン

```bash
git clone https://github.com/w0k-sr/docker.git
cd docker/sudoku
```
### 2. Dockerでアプリケーションをビルド＆起動
```bash
# ビルド
docker-compose up --build
# 起動
docker exec -it my_app bash
/go/src/app> cd myapp
/go/src/app/myapp> bee run
```
## 🖥️ 画面イメージ
(今後アップデート)

## 💡 工夫ポイント
- **MVC設計（Beego）**によりコードの責務を分離し、保守性を向上
- Docker化により、環境依存を排除し誰でもすぐに動作確認可能
- GitHub上でソースコードを公開し、技術ポートフォリオとしても活用

## 📚 参考文献
- [BeegoによるWebアプリケーション開発① on docker](https://qiita.com/uchidash456/items/f940a0eba24c53529bb4)