# go-docker-clock

[FROM scratchから始める軽量Docker image for Go \- Qiita](https://qiita.com/Saint1991/items/dcd6a92e5074bd10f75a)
にあったコードをそのままコピペして、サポートのスクリプトをつけたもの。

時間を文字列で返すタイムサーバもどき。

実行例:
```
$ ./go-docker-clock &
$ curl http://localhost:8080/time
2021-09-07 01:47:04 UTC
```

Docker imageは2.16MB
```
$ docker images go-docker-clock
REPOSITORY        TAG       IMAGE ID       CREATED          SIZE
go-docker-clock   latest    e9161e29eb40   16 minutes ago   2.16MB
```


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
