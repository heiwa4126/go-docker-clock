# go-docker-clock

[FROM scratchから始める軽量Docker image for Go \- Qiita](https://qiita.com/Saint1991/items/dcd6a92e5074bd10f75a)
にあったコードをそのままコピペして、
サポートのスクリプトをつけたもの。

# golangレベルでの開発

```sh
go build
./go-docker-clock &
./test.sh
kill %1   # ここは要アレンジ
```

# docker imageの作成

```sh
./build-go.sh
./build-docker-image.sh
./run-docker.sh
./test.sh
./stop-docker.sh
```
