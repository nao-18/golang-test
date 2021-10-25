# dockerによるgolang開発環境(air含む)

## ビルドコマンド
~~~bash
docker-compose buid
~~~

## 依存関係解決コマンド
~~~bash
docker-compose run --rm app go mod tidy
~~~

## 起動コマンド
~~~bash
docker-compose up -d
~~~

# テスト
## テスト
~~~bash
docker-compose run --rm app go test
~~~

## カバレッジ、テスト詳細
~~~bash
docker-compose run --rm app go test -v -cover
~~~

## カバレッジプロファイル出力
~~~bash
docker-compose run --rm app go test -cover ./... -coverprofile=cover.out
~~~

## cover.html作成
~~~bash
docker-compose run --rm app go tool cover -html=cover.out -o cover.html
~~~