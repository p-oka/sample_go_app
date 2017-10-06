# sample_go_app
golangの開発環境をDockerで構築するためのサンプル

## 要件
- ホストマシン(Mac)からポートフォワーディングできること
- Golangで書いたWebAppが別コンテナとして立ち上げたMySQLに疎通できること

## 検証パターン
- [x] docker-composeを使う
- [x] minikubeを使う
