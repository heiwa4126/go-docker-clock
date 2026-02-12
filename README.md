# go-docker-clock

[FROM scratch から始める軽量 Docker image for Go \- Qiita](https://qiita.com/Saint1991/items/dcd6a92e5074bd10f75a)
にあったコードをそのままコピペして、lint が通るようにしてサポートのスクリプトをつけたもの。

時間を文字列で返すタイムサーバもどき。

## Docker image を GitHub Package に発行

semver の tag をつけて GitHub に push することで、
GitHub Package にイメージをパブリッシュします。

発行先:
[Package go-docker-clock](https://github.com/heiwa4126/go-docker-clock/pkgs/container/go-docker-clock)

使い方は

```sh
docker run --rm -d -p 8080:8080 ghcr.io/heiwa4126/go-docker-clock:latest
curl "localhost:8080/time?tz=Asia/Tokyo"
```

**重要**: push の前に `task` (引数なし)を実行しよう。linter が動きます

## 開発準備

linter などは CI/CD できるよう
[aquaproj/aqua](https://github.com/aquaproj/aqua)
で管理しています。

プロジェクトルートで

```sh
aqua i
```

でインストール。
(Go 自体は入ってないです。必要なら `aqua i golang/go`)

## 開発

```sh
# 普通に実行
task start
./smoke-test.sh
task stop

# lintとauditとformat
task audit
task test
task lint
task fmt
task   # 全部まとめて (default action)

# ビルドして実行
task build-go
task start-go
./smoke-test.sh
task kill

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
