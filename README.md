# go-docker-clock

[FROM scratch から始める軽量 Docker image for Go \- Qiita](https://qiita.com/Saint1991/items/dcd6a92e5074bd10f75a)
にあったコードをそのままコピペして、lintが通るようにしてサポートのスクリプトをつけたもの。

時間を文字列で返すタイムサーバもどき。

## 準備

linter などは CI/CD できるよう
[aquaproj/aqua](https://github.com/aquaproj/aqua)
で管理しています。

プロジェクトルートで

```sh
aqua i
```

でインストール

## 開発

```sh
# 普通に実行
go run . &
pid=$!
./smoke-test.sh
kill "$pid"

# lintとauditとformat
task   # default action

# ビルドして実行
task build-go
./go-docker-clock&
pid=$!
./smoke-test.sh
kill "$pid"

# コンテナ作って実行
task build-docker  # `task build-go` してから実行
task start-docker
./smoke-test.sh
task stop-docker
```

## 出力例

```console
$ ./go-docker-clock &
$ curl http://localhost:8080/time
2021-09-07 01:47:04 UTC
```

## バイナリも Docker image も 2MB 強

```console
$ docker images go-docker-clock
REPOSITORY        TAG       IMAGE ID       CREATED          SIZE
go-docker-clock   latest    e9161e29eb40   16 minutes ago   2.16MB
```

# おまけ - docker imageをローカルファイルにsave/load

セーブ

```sh
docker save go-docker-clock | gzip -9 >  go-docker-clock.tar.gz
```

ロード

```sh
zcat go-docker-clock.tar.gz | docker load
```
