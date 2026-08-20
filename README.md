# go-docker-clock

[FROM scratch から始める軽量 Docker image for Go \- Qiita](https://qiita.com/Saint1991/items/dcd6a92e5074bd10f75a)
にあったコードをそのままコピペして、lint が通るようにしてサポートのスクリプトをつけたもの。

時間を文字列で返すタイムサーバもどき。

Go 1.25 使用

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
task build-docker
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

→ いろいろ考えて UPX やめました。それでも 8MB

## zoneinfo.zip 式をやめる

Dockerfile の

```dockerfile
FROM scratch

ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO=/zoneinfo.zip
# ...
```

をやめて、ビルド時に `-tags timetzdata` をつける式にした。Go 1.15 から使えるらしい。[Go 1.15 Release Notes - The Go Programming Language](https://go.dev/doc/go1.15#time_tzdata)

※ もう 1 つの方法は `import _ "time/tzdata"`

もう 1 つ

```sh
task build-docker-distroless
```

も作ってみた。

## おまけ - docker imageをローカルファイルにsave/load

セーブ

```sh
docker save go-docker-clock | gzip -9 >  go-docker-clock.tar.gz
```

ロード

```sh
zcat go-docker-clock.tar.gz | docker load
```
