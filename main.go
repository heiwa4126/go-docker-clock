package main

import (
	"net/http"
	"time"
)

const layout = "2006-01-02 15:04:05 MST"

func main() {
	http.HandleFunc("/time",
		func(writer http.ResponseWriter, request *http.Request) {
			l := request.URL.Query().Get("tz")
			location, err := time.LoadLocation(l)
			if err != nil {
				panic(err)
			}
			if _, err := writer.Write([]byte(time.Now().In(location).Format(layout))); err != nil {
				// エラーが発生しても既にレスポンスの書き込みが始まっている可能性があるため、
				// ログを記録して処理を終了する（サーバーは停止しない）
				// 本番環境では適切なロギングライブラリを使用する
				return
			}
		})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
