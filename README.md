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

# docker imageの作成とテスト

```sh
./build-go.sh　# 要upx
./build-docker-image.sh
./run-docker.sh
./test.sh
./stop-docker.sh
```

# おまけ - docker imageをローカルファイルにsave/load

セーブ
```sh
docker save go-docker-clock | gzip -9 >  go-docker-clock.tar.gz
```

ロード
```
zcat go-docker-clock.tar.gz | docker load
```

## TODO

export/importも試す
